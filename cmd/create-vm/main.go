package main

import (
	"net/http"

	goarg "github.com/alexflint/go-arg"
	. "github.com/kubevirt/kubevirt-tekton-tasks/modules/create-vm/pkg/constants"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/create-vm/pkg/utils/parse"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/create-vm/pkg/vmcreator"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/exit"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/log"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/output"
	res "github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/results"
	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/zerrors"
	"go.uber.org/zap"
	kubevirtv1 "kubevirt.io/api/core/v1"
)

func main() {
	defer exit.HandleExit()

	cliOptions := &parse.CLIOptions{}
	goarg.MustParse(cliOptions)

	logger := log.InitLogger(cliOptions.GetDebugLevel())
	defer logger.Sync()

	if err := cliOptions.Init(); err != nil {
		exit.ExitOrDieFromError(InvalidCLIInputExitCode, err)
	}
	log.Logger().Debug("parsed arguments", zap.Reflect("cliOptions", cliOptions))

	vmCreator, err := vmcreator.NewVMCreator(cliOptions)
	if err != nil {
		exit.ExitOrDieFromError(GenericExitCode, err)
	}

	vm, err := vmCreator.CreateVM()
	if err != nil {
		exit.ExitOrDieFromError(CreateVMErrorExitCode, err,
			zerrors.IsStatusError(err, http.StatusNotFound, http.StatusConflict, http.StatusUnprocessableEntity),
		)
	}

	if cliOptions.GetStartVMFlag() &&
		(vm.Spec.RunStrategy == nil || *vm.Spec.RunStrategy != kubevirtv1.RunStrategyAlways) &&
		(vm.Spec.Running == nil || !*vm.Spec.Running) {
		if err := vmCreator.StartVM(vm.Namespace, vm.Name); err != nil {
			exit.ExitFromError(StartVMErrorExitCode, err)
		}
	}

	results := map[string]string{
		NameResultName:      vm.Name,
		NamespaceResultName: vm.Namespace,
	}

	log.Logger().Debug("recording results", zap.Reflect("results", results))
	if err := res.RecordResults(results); err != nil {
		exit.ExitOrDieFromError(WriteResultsExitCode, err)
	}

	output.PrettyPrint(vm, cliOptions.Output)
}
