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
			in.(*Activity).DeepCopyInto(out.(*Activity))
			return nil
		}, InType: reflect.TypeOf(&Activity{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActivityList).DeepCopyInto(out.(*ActivityList))
			return nil
		}, InType: reflect.TypeOf(&ActivityList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActivitySpec).DeepCopyInto(out.(*ActivitySpec))
			return nil
		}, InType: reflect.TypeOf(&ActivitySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActivityStage).DeepCopyInto(out.(*ActivityStage))
			return nil
		}, InType: reflect.TypeOf(&ActivityStage{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActivityStatus).DeepCopyInto(out.(*ActivityStatus))
			return nil
		}, InType: reflect.TypeOf(&ActivityStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActivityStep).DeepCopyInto(out.(*ActivityStep))
			return nil
		}, InType: reflect.TypeOf(&ActivityStep{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CIService).DeepCopyInto(out.(*CIService))
			return nil
		}, InType: reflect.TypeOf(&CIService{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CronTrigger).DeepCopyInto(out.(*CronTrigger))
			return nil
		}, InType: reflect.TypeOf(&CronTrigger{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitAccount).DeepCopyInto(out.(*GitAccount))
			return nil
		}, InType: reflect.TypeOf(&GitAccount{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GitAccountList).DeepCopyInto(out.(*GitAccountList))
			return nil
		}, InType: reflect.TypeOf(&GitAccountList{})},
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
			in.(*GitRepoCacheList).DeepCopyInto(out.(*GitRepoCacheList))
			return nil
		}, InType: reflect.TypeOf(&GitRepoCacheList{})},
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
			in.(*PipelineConditions).DeepCopyInto(out.(*PipelineConditions))
			return nil
		}, InType: reflect.TypeOf(&PipelineConditions{})},
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
			in.(*SCMSetting).DeepCopyInto(out.(*SCMSetting))
			return nil
		}, InType: reflect.TypeOf(&SCMSetting{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SCMSettingList).DeepCopyInto(out.(*SCMSettingList))
			return nil
		}, InType: reflect.TypeOf(&SCMSettingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SCMSettingSpec).DeepCopyInto(out.(*SCMSettingSpec))
			return nil
		}, InType: reflect.TypeOf(&SCMSettingSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SCMSettingStatus).DeepCopyInto(out.(*SCMSettingStatus))
			return nil
		}, InType: reflect.TypeOf(&SCMSettingStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Stage).DeepCopyInto(out.(*Stage))
			return nil
		}, InType: reflect.TypeOf(&Stage{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Step).DeepCopyInto(out.(*Step))
			return nil
		}, InType: reflect.TypeOf(&Step{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Activity) DeepCopyInto(out *Activity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ActivitySpec.DeepCopyInto(&out.ActivitySpec)
	out.ActivityStatus = in.ActivityStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Activity.
func (in *Activity) DeepCopy() *Activity {
	if in == nil {
		return nil
	}
	out := new(Activity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Activity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityList) DeepCopyInto(out *ActivityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Activity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityList.
func (in *ActivityList) DeepCopy() *ActivityList {
	if in == nil {
		return nil
	}
	out := new(ActivityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActivityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivitySpec) DeepCopyInto(out *ActivitySpec) {
	*out = *in
	in.Pipeline.DeepCopyInto(&out.Pipeline)
	if in.ActivityStages != nil {
		in, out := &in.ActivityStages, &out.ActivityStages
		*out = make([]ActivityStage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivitySpec.
func (in *ActivitySpec) DeepCopy() *ActivitySpec {
	if in == nil {
		return nil
	}
	out := new(ActivitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityStage) DeepCopyInto(out *ActivityStage) {
	*out = *in
	if in.Approvers != nil {
		in, out := &in.Approvers, &out.Approvers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActivitySteps != nil {
		in, out := &in.ActivitySteps, &out.ActivitySteps
		*out = make([]ActivityStep, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityStage.
func (in *ActivityStage) DeepCopy() *ActivityStage {
	if in == nil {
		return nil
	}
	out := new(ActivityStage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityStatus) DeepCopyInto(out *ActivityStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityStatus.
func (in *ActivityStatus) DeepCopy() *ActivityStatus {
	if in == nil {
		return nil
	}
	out := new(ActivityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityStep) DeepCopyInto(out *ActivityStep) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityStep.
func (in *ActivityStep) DeepCopy() *ActivityStep {
	if in == nil {
		return nil
	}
	out := new(ActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CIService) DeepCopyInto(out *CIService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIService.
func (in *CIService) DeepCopy() *CIService {
	if in == nil {
		return nil
	}
	out := new(CIService)
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
func (in *GitAccount) DeepCopyInto(out *GitAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.GitAccountSpec = in.GitAccountSpec
	out.GitAccountStatus = in.GitAccountStatus
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
func (in *GitAccountList) DeepCopyInto(out *GitAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitAccountList.
func (in *GitAccountList) DeepCopy() *GitAccountList {
	if in == nil {
		return nil
	}
	out := new(GitAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitAccountList) DeepCopyObject() runtime.Object {
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
func (in *GitRepoCacheList) DeepCopyInto(out *GitRepoCacheList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitRepoCache, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepoCacheList.
func (in *GitRepoCacheList) DeepCopy() *GitRepoCacheList {
	if in == nil {
		return nil
	}
	out := new(GitRepoCacheList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitRepoCacheList) DeepCopyObject() runtime.Object {
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
	in.PipelineSpec.DeepCopyInto(&out.PipelineSpec)
	out.PipelineStatus = in.PipelineStatus
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
func (in *PipelineConditions) DeepCopyInto(out *PipelineConditions) {
	*out = *in
	if in.All != nil {
		in, out := &in.All, &out.All
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Any != nil {
		in, out := &in.Any, &out.Any
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineConditions.
func (in *PipelineConditions) DeepCopy() *PipelineConditions {
	if in == nil {
		return nil
	}
	out := new(PipelineConditions)
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
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
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
func (in *SCMSetting) DeepCopyInto(out *SCMSetting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCMSetting.
func (in *SCMSetting) DeepCopy() *SCMSetting {
	if in == nil {
		return nil
	}
	out := new(SCMSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SCMSetting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCMSettingList) DeepCopyInto(out *SCMSettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SCMSetting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCMSettingList.
func (in *SCMSettingList) DeepCopy() *SCMSettingList {
	if in == nil {
		return nil
	}
	out := new(SCMSettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SCMSettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCMSettingSpec) DeepCopyInto(out *SCMSettingSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCMSettingSpec.
func (in *SCMSettingSpec) DeepCopy() *SCMSettingSpec {
	if in == nil {
		return nil
	}
	out := new(SCMSettingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCMSettingStatus) DeepCopyInto(out *SCMSettingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCMSettingStatus.
func (in *SCMSettingStatus) DeepCopy() *SCMSettingStatus {
	if in == nil {
		return nil
	}
	out := new(SCMSettingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stage) DeepCopyInto(out *Stage) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	if in.Approvers != nil {
		in, out := &in.Approvers, &out.Approvers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
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
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]CIService, len(*in))
		copy(*out, *in)
	}
	if in.ServiceSelector != nil {
		in, out := &in.ServiceSelector, &out.ServiceSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
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
