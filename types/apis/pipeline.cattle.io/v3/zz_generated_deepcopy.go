package v3

import (
	reflect "reflect"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BuildImageStepConfig).DeepCopyInto(out.(*BuildImageStepConfig))
			return nil
		}, InType: reflect.TypeOf(&BuildImageStepConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CronTrigger).DeepCopyInto(out.(*CronTrigger))
			return nil
		}, InType: reflect.TypeOf(&CronTrigger{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GibhubConfig).DeepCopyInto(out.(*GibhubConfig))
			return nil
		}, InType: reflect.TypeOf(&GibhubConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitAccount).DeepCopyInto(out.(*GitAccount))
			return nil
		}, InType: reflect.TypeOf(&GitAccount{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitAccountSpec).DeepCopyInto(out.(*GitAccountSpec))
			return nil
		}, InType: reflect.TypeOf(&GitAccountSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitAccountStatus).DeepCopyInto(out.(*GitAccountStatus))
			return nil
		}, InType: reflect.TypeOf(&GitAccountStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitRepoCache).DeepCopyInto(out.(*GitRepoCache))
			return nil
		}, InType: reflect.TypeOf(&GitRepoCache{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitRepoCacheSpec).DeepCopyInto(out.(*GitRepoCacheSpec))
			return nil
		}, InType: reflect.TypeOf(&GitRepoCacheSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitRepoCacheStatus).DeepCopyInto(out.(*GitRepoCacheStatus))
			return nil
		}, InType: reflect.TypeOf(&GitRepoCacheStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitRepository).DeepCopyInto(out.(*GitRepository))
			return nil
		}, InType: reflect.TypeOf(&GitRepository{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Pipeline).DeepCopyInto(out.(*Pipeline))
			return nil
		}, InType: reflect.TypeOf(&Pipeline{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineConfig).DeepCopyInto(out.(*PipelineConfig))
			return nil
		}, InType: reflect.TypeOf(&PipelineConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineHistory).DeepCopyInto(out.(*PipelineHistory))
			return nil
		}, InType: reflect.TypeOf(&PipelineHistory{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineHistoryList).DeepCopyInto(out.(*PipelineHistoryList))
			return nil
		}, InType: reflect.TypeOf(&PipelineHistoryList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineHistorySpec).DeepCopyInto(out.(*PipelineHistorySpec))
			return nil
		}, InType: reflect.TypeOf(&PipelineHistorySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineHistoryStatus).DeepCopyInto(out.(*PipelineHistoryStatus))
			return nil
		}, InType: reflect.TypeOf(&PipelineHistoryStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineList).DeepCopyInto(out.(*PipelineList))
			return nil
		}, InType: reflect.TypeOf(&PipelineList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineSpec).DeepCopyInto(out.(*PipelineSpec))
			return nil
		}, InType: reflect.TypeOf(&PipelineSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PipelineStatus).DeepCopyInto(out.(*PipelineStatus))
			return nil
		}, InType: reflect.TypeOf(&PipelineStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RunScriptStepConfig).DeepCopyInto(out.(*RunScriptStepConfig))
			return nil
		}, InType: reflect.TypeOf(&RunScriptStepConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SourceCodeStepConfig).DeepCopyInto(out.(*SourceCodeStepConfig))
			return nil
		}, InType: reflect.TypeOf(&SourceCodeStepConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Stage).DeepCopyInto(out.(*Stage))
			return nil
		}, InType: reflect.TypeOf(&Stage{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StageStatus).DeepCopyInto(out.(*StageStatus))
			return nil
		}, InType: reflect.TypeOf(&StageStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Step).DeepCopyInto(out.(*Step))
			return nil
		}, InType: reflect.TypeOf(&Step{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StepStatus).DeepCopyInto(out.(*StepStatus))
			return nil
		}, InType: reflect.TypeOf(&StepStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildImageStepConfig) DeepCopyInto(out *BuildImageStepConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildImageStepConfig.
func (in *BuildImageStepConfig) DeepCopy() *BuildImageStepConfig {
	if in == nil {
		return nil
	}
	out := new(BuildImageStepConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronTrigger) DeepCopyInto(out *CronTrigger) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronTrigger.
func (in *CronTrigger) DeepCopy() *CronTrigger {
	if in == nil {
		return nil
	}
	out := new(CronTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GibhubConfig) DeepCopyInto(out *GibhubConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GibhubConfig.
func (in *GibhubConfig) DeepCopy() *GibhubConfig {
	if in == nil {
		return nil
	}
	out := new(GibhubConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitAccount) DeepCopyInto(out *GitAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitAccount.
func (in *GitAccount) DeepCopy() *GitAccount {
	if in == nil {
		return nil
	}
	out := new(GitAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitAccountSpec) DeepCopyInto(out *GitAccountSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitAccountSpec.
func (in *GitAccountSpec) DeepCopy() *GitAccountSpec {
	if in == nil {
		return nil
	}
	out := new(GitAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitAccountStatus) DeepCopyInto(out *GitAccountStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitAccountStatus.
func (in *GitAccountStatus) DeepCopy() *GitAccountStatus {
	if in == nil {
		return nil
	}
	out := new(GitAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepoCache) DeepCopyInto(out *GitRepoCache) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepoCache.
func (in *GitRepoCache) DeepCopy() *GitRepoCache {
	if in == nil {
		return nil
	}
	out := new(GitRepoCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitRepoCache) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepoCacheSpec) DeepCopyInto(out *GitRepoCacheSpec) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]GitRepository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepoCacheSpec.
func (in *GitRepoCacheSpec) DeepCopy() *GitRepoCacheSpec {
	if in == nil {
		return nil
	}
	out := new(GitRepoCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepoCacheStatus) DeepCopyInto(out *GitRepoCacheStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepoCacheStatus.
func (in *GitRepoCacheStatus) DeepCopy() *GitRepoCacheStatus {
	if in == nil {
		return nil
	}
	out := new(GitRepoCacheStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepository) DeepCopyInto(out *GitRepository) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepository.
func (in *GitRepository) DeepCopy() *GitRepository {
	if in == nil {
		return nil
	}
	out := new(GitRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineConfig) DeepCopyInto(out *PipelineConfig) {
	*out = *in
	if in.GibhubConfig != nil {
		in, out := &in.GibhubConfig, &out.GibhubConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(GibhubConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineConfig.
func (in *PipelineConfig) DeepCopy() *PipelineConfig {
	if in == nil {
		return nil
	}
	out := new(PipelineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineHistory) DeepCopyInto(out *PipelineHistory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineHistory.
func (in *PipelineHistory) DeepCopy() *PipelineHistory {
	if in == nil {
		return nil
	}
	out := new(PipelineHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineHistory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineHistoryList) DeepCopyInto(out *PipelineHistoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PipelineHistory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineHistoryList.
func (in *PipelineHistoryList) DeepCopy() *PipelineHistoryList {
	if in == nil {
		return nil
	}
	out := new(PipelineHistoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineHistoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineHistorySpec) DeepCopyInto(out *PipelineHistorySpec) {
	*out = *in
	in.Pipeline.DeepCopyInto(&out.Pipeline)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineHistorySpec.
func (in *PipelineHistorySpec) DeepCopy() *PipelineHistorySpec {
	if in == nil {
		return nil
	}
	out := new(PipelineHistorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineHistoryStatus) DeepCopyInto(out *PipelineHistoryStatus) {
	*out = *in
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StageStatus != nil {
		in, out := &in.StageStatus, &out.StageStatus
		*out = make([]StageStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineHistoryStatus.
func (in *PipelineHistoryStatus) DeepCopy() *PipelineHistoryStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineHistoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineList) DeepCopyInto(out *PipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineList.
func (in *PipelineList) DeepCopy() *PipelineList {
	if in == nil {
		return nil
	}
	out := new(PipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	out.CronTrigger = in.CronTrigger
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]Stage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunScriptStepConfig) DeepCopyInto(out *RunScriptStepConfig) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunScriptStepConfig.
func (in *RunScriptStepConfig) DeepCopy() *RunScriptStepConfig {
	if in == nil {
		return nil
	}
	out := new(RunScriptStepConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceCodeStepConfig) DeepCopyInto(out *SourceCodeStepConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceCodeStepConfig.
func (in *SourceCodeStepConfig) DeepCopy() *SourceCodeStepConfig {
	if in == nil {
		return nil
	}
	out := new(SourceCodeStepConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stage) DeepCopyInto(out *Stage) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]Step, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stage.
func (in *Stage) DeepCopy() *Stage {
	if in == nil {
		return nil
	}
	out := new(Stage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageStatus) DeepCopyInto(out *StageStatus) {
	*out = *in
	if in.StepStatus != nil {
		in, out := &in.StepStatus, &out.StepStatus
		*out = make([]StepStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageStatus.
func (in *StageStatus) DeepCopy() *StageStatus {
	if in == nil {
		return nil
	}
	out := new(StageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	if in.SourceCodeStepConfig != nil {
		in, out := &in.SourceCodeStepConfig, &out.SourceCodeStepConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(SourceCodeStepConfig)
			**out = **in
		}
	}
	if in.RunScriptStepConfig != nil {
		in, out := &in.RunScriptStepConfig, &out.RunScriptStepConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(RunScriptStepConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.BuildImageStepConfig != nil {
		in, out := &in.BuildImageStepConfig, &out.BuildImageStepConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(BuildImageStepConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Step.
func (in *Step) DeepCopy() *Step {
	if in == nil {
		return nil
	}
	out := new(Step)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepStatus) DeepCopyInto(out *StepStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepStatus.
func (in *StepStatus) DeepCopy() *StepStatus {
	if in == nil {
		return nil
	}
	out := new(StepStatus)
	in.DeepCopyInto(out)
	return out
}
