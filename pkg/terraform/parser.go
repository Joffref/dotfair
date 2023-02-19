package terraform

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
	"os"
)

const (
	planFile = "plan.tfplan"
)

// Plan executes a dry run of terraform and returns the plan output
func Plan(ctx context.Context, workspace string) (*tfjson.Plan, error) {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.3.9")),
	}
	execPath, err := installer.Install(ctx)
	if err != nil {
		return nil, err
	}
	tf, err := tfexec.NewTerraform(workspace, execPath)
	if err != nil {
		return nil, err
	}
	err = tf.Init(ctx, tfexec.Upgrade(true))
	if err != nil {
		return nil, err
	}
	ok, err := tf.Plan(ctx, tfexec.Out(planFile))
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("terraform plan failed")
	}
	plan, err := tf.ShowPlanFile(ctx, planFile)
	if err != nil {
		return nil, err
	}
	err = os.Remove(fmt.Sprintf("%s/%s", workspace, planFile))
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func CurrentState(ctx context.Context, workspace string) (*tfjson.State, error) {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.3.9")),
	}
	execPath, err := installer.Install(ctx)
	if err != nil {
		return nil, err
	}
	tf, err := tfexec.NewTerraform(workspace, execPath)
	if err != nil {
		return nil, err
	}
	err = tf.Init(ctx, tfexec.Upgrade(true))
	if err != nil {
		return nil, err
	}
	state, err := tf.Show(ctx)
	if err != nil {
		return nil, err
	}
	return state, nil
}
