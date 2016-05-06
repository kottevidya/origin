// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	image_api "github.com/openshift/origin/pkg/image/api"
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
	intstr "k8s.io/kubernetes/pkg/util/intstr"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_api_CustomDeploymentStrategyParams,
		DeepCopy_api_DeploymentCause,
		DeepCopy_api_DeploymentCauseImageTrigger,
		DeepCopy_api_DeploymentConfig,
		DeepCopy_api_DeploymentConfigList,
		DeepCopy_api_DeploymentConfigRollback,
		DeepCopy_api_DeploymentConfigRollbackSpec,
		DeepCopy_api_DeploymentConfigSpec,
		DeepCopy_api_DeploymentConfigStatus,
		DeepCopy_api_DeploymentDetails,
		DeepCopy_api_DeploymentLog,
		DeepCopy_api_DeploymentLogOptions,
		DeepCopy_api_DeploymentStrategy,
		DeepCopy_api_DeploymentTriggerImageChangeParams,
		DeepCopy_api_DeploymentTriggerPolicy,
		DeepCopy_api_ExecNewPodHook,
		DeepCopy_api_LifecycleHook,
		DeepCopy_api_RecreateDeploymentStrategyParams,
		DeepCopy_api_RollingDeploymentStrategyParams,
		DeepCopy_api_TagImageHook,
		DeepCopy_api_TemplateImage,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_api_CustomDeploymentStrategyParams(in CustomDeploymentStrategyParams, out *CustomDeploymentStrategyParams, c *conversion.Cloner) error {
	out.Image = in.Image
	if in.Environment != nil {
		in, out := in.Environment, &out.Environment
		*out = make([]api.EnvVar, len(in))
		for i := range in {
			if err := api.DeepCopy_api_EnvVar(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Environment = nil
	}
	if in.Command != nil {
		in, out := in.Command, &out.Command
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Command = nil
	}
	return nil
}

func DeepCopy_api_DeploymentCause(in DeploymentCause, out *DeploymentCause, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.ImageTrigger != nil {
		in, out := in.ImageTrigger, &out.ImageTrigger
		*out = new(DeploymentCauseImageTrigger)
		if err := DeepCopy_api_DeploymentCauseImageTrigger(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.ImageTrigger = nil
	}
	return nil
}

func DeepCopy_api_DeploymentCauseImageTrigger(in DeploymentCauseImageTrigger, out *DeploymentCauseImageTrigger, c *conversion.Cloner) error {
	if err := api.DeepCopy_api_ObjectReference(in.From, &out.From, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_DeploymentConfig(in DeploymentConfig, out *DeploymentConfig, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api.DeepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_api_DeploymentConfigSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_api_DeploymentConfigStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_DeploymentConfigList(in DeploymentConfigList, out *DeploymentConfigList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]DeploymentConfig, len(in))
		for i := range in {
			if err := DeepCopy_api_DeploymentConfig(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_api_DeploymentConfigRollback(in DeploymentConfigRollback, out *DeploymentConfigRollback, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_api_DeploymentConfigRollbackSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_DeploymentConfigRollbackSpec(in DeploymentConfigRollbackSpec, out *DeploymentConfigRollbackSpec, c *conversion.Cloner) error {
	if err := api.DeepCopy_api_ObjectReference(in.From, &out.From, c); err != nil {
		return err
	}
	out.IncludeTriggers = in.IncludeTriggers
	out.IncludeTemplate = in.IncludeTemplate
	out.IncludeReplicationMeta = in.IncludeReplicationMeta
	out.IncludeStrategy = in.IncludeStrategy
	return nil
}

func DeepCopy_api_DeploymentConfigSpec(in DeploymentConfigSpec, out *DeploymentConfigSpec, c *conversion.Cloner) error {
	if err := DeepCopy_api_DeploymentStrategy(in.Strategy, &out.Strategy, c); err != nil {
		return err
	}
	if in.Triggers != nil {
		in, out := in.Triggers, &out.Triggers
		*out = make([]DeploymentTriggerPolicy, len(in))
		for i := range in {
			if err := DeepCopy_api_DeploymentTriggerPolicy(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	out.Replicas = in.Replicas
	out.Test = in.Test
	if in.Selector != nil {
		in, out := in.Selector, &out.Selector
		*out = make(map[string]string)
		for key, val := range in {
			(*out)[key] = val
		}
	} else {
		out.Selector = nil
	}
	if in.Template != nil {
		in, out := in.Template, &out.Template
		*out = new(api.PodTemplateSpec)
		if err := api.DeepCopy_api_PodTemplateSpec(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Template = nil
	}
	return nil
}

func DeepCopy_api_DeploymentConfigStatus(in DeploymentConfigStatus, out *DeploymentConfigStatus, c *conversion.Cloner) error {
	out.LatestVersion = in.LatestVersion
	if in.Details != nil {
		in, out := in.Details, &out.Details
		*out = new(DeploymentDetails)
		if err := DeepCopy_api_DeploymentDetails(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Details = nil
	}
	return nil
}

func DeepCopy_api_DeploymentDetails(in DeploymentDetails, out *DeploymentDetails, c *conversion.Cloner) error {
	out.Message = in.Message
	if in.Causes != nil {
		in, out := in.Causes, &out.Causes
		*out = make([]DeploymentCause, len(in))
		for i := range in {
			if err := DeepCopy_api_DeploymentCause(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Causes = nil
	}
	return nil
}

func DeepCopy_api_DeploymentLog(in DeploymentLog, out *DeploymentLog, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_DeploymentLogOptions(in DeploymentLogOptions, out *DeploymentLogOptions, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Container = in.Container
	out.Follow = in.Follow
	out.Previous = in.Previous
	if in.SinceSeconds != nil {
		in, out := in.SinceSeconds, &out.SinceSeconds
		*out = new(int64)
		**out = *in
	} else {
		out.SinceSeconds = nil
	}
	if in.SinceTime != nil {
		in, out := in.SinceTime, &out.SinceTime
		*out = new(unversioned.Time)
		if err := unversioned.DeepCopy_unversioned_Time(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.SinceTime = nil
	}
	out.Timestamps = in.Timestamps
	if in.TailLines != nil {
		in, out := in.TailLines, &out.TailLines
		*out = new(int64)
		**out = *in
	} else {
		out.TailLines = nil
	}
	if in.LimitBytes != nil {
		in, out := in.LimitBytes, &out.LimitBytes
		*out = new(int64)
		**out = *in
	} else {
		out.LimitBytes = nil
	}
	out.NoWait = in.NoWait
	if in.Version != nil {
		in, out := in.Version, &out.Version
		*out = new(int64)
		**out = *in
	} else {
		out.Version = nil
	}
	return nil
}

func DeepCopy_api_DeploymentStrategy(in DeploymentStrategy, out *DeploymentStrategy, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.CustomParams != nil {
		in, out := in.CustomParams, &out.CustomParams
		*out = new(CustomDeploymentStrategyParams)
		if err := DeepCopy_api_CustomDeploymentStrategyParams(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.CustomParams = nil
	}
	if in.RecreateParams != nil {
		in, out := in.RecreateParams, &out.RecreateParams
		*out = new(RecreateDeploymentStrategyParams)
		if err := DeepCopy_api_RecreateDeploymentStrategyParams(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.RecreateParams = nil
	}
	if in.RollingParams != nil {
		in, out := in.RollingParams, &out.RollingParams
		*out = new(RollingDeploymentStrategyParams)
		if err := DeepCopy_api_RollingDeploymentStrategyParams(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.RollingParams = nil
	}
	if err := api.DeepCopy_api_ResourceRequirements(in.Resources, &out.Resources, c); err != nil {
		return err
	}
	if in.Labels != nil {
		in, out := in.Labels, &out.Labels
		*out = make(map[string]string)
		for key, val := range in {
			(*out)[key] = val
		}
	} else {
		out.Labels = nil
	}
	if in.Annotations != nil {
		in, out := in.Annotations, &out.Annotations
		*out = make(map[string]string)
		for key, val := range in {
			(*out)[key] = val
		}
	} else {
		out.Annotations = nil
	}
	return nil
}

func DeepCopy_api_DeploymentTriggerImageChangeParams(in DeploymentTriggerImageChangeParams, out *DeploymentTriggerImageChangeParams, c *conversion.Cloner) error {
	out.Automatic = in.Automatic
	if in.ContainerNames != nil {
		in, out := in.ContainerNames, &out.ContainerNames
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.ContainerNames = nil
	}
	if err := api.DeepCopy_api_ObjectReference(in.From, &out.From, c); err != nil {
		return err
	}
	out.LastTriggeredImage = in.LastTriggeredImage
	return nil
}

func DeepCopy_api_DeploymentTriggerPolicy(in DeploymentTriggerPolicy, out *DeploymentTriggerPolicy, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.ImageChangeParams != nil {
		in, out := in.ImageChangeParams, &out.ImageChangeParams
		*out = new(DeploymentTriggerImageChangeParams)
		if err := DeepCopy_api_DeploymentTriggerImageChangeParams(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.ImageChangeParams = nil
	}
	return nil
}

func DeepCopy_api_ExecNewPodHook(in ExecNewPodHook, out *ExecNewPodHook, c *conversion.Cloner) error {
	if in.Command != nil {
		in, out := in.Command, &out.Command
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Command = nil
	}
	if in.Env != nil {
		in, out := in.Env, &out.Env
		*out = make([]api.EnvVar, len(in))
		for i := range in {
			if err := api.DeepCopy_api_EnvVar(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Env = nil
	}
	out.ContainerName = in.ContainerName
	if in.Volumes != nil {
		in, out := in.Volumes, &out.Volumes
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Volumes = nil
	}
	return nil
}

func DeepCopy_api_LifecycleHook(in LifecycleHook, out *LifecycleHook, c *conversion.Cloner) error {
	out.FailurePolicy = in.FailurePolicy
	if in.ExecNewPod != nil {
		in, out := in.ExecNewPod, &out.ExecNewPod
		*out = new(ExecNewPodHook)
		if err := DeepCopy_api_ExecNewPodHook(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.ExecNewPod = nil
	}
	if in.TagImages != nil {
		in, out := in.TagImages, &out.TagImages
		*out = make([]TagImageHook, len(in))
		for i := range in {
			if err := DeepCopy_api_TagImageHook(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.TagImages = nil
	}
	return nil
}

func DeepCopy_api_RecreateDeploymentStrategyParams(in RecreateDeploymentStrategyParams, out *RecreateDeploymentStrategyParams, c *conversion.Cloner) error {
	if in.TimeoutSeconds != nil {
		in, out := in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int64)
		**out = *in
	} else {
		out.TimeoutSeconds = nil
	}
	if in.Pre != nil {
		in, out := in.Pre, &out.Pre
		*out = new(LifecycleHook)
		if err := DeepCopy_api_LifecycleHook(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Mid != nil {
		in, out := in.Mid, &out.Mid
		*out = new(LifecycleHook)
		if err := DeepCopy_api_LifecycleHook(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Mid = nil
	}
	if in.Post != nil {
		in, out := in.Post, &out.Post
		*out = new(LifecycleHook)
		if err := DeepCopy_api_LifecycleHook(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}

func DeepCopy_api_RollingDeploymentStrategyParams(in RollingDeploymentStrategyParams, out *RollingDeploymentStrategyParams, c *conversion.Cloner) error {
	if in.UpdatePeriodSeconds != nil {
		in, out := in.UpdatePeriodSeconds, &out.UpdatePeriodSeconds
		*out = new(int64)
		**out = *in
	} else {
		out.UpdatePeriodSeconds = nil
	}
	if in.IntervalSeconds != nil {
		in, out := in.IntervalSeconds, &out.IntervalSeconds
		*out = new(int64)
		**out = *in
	} else {
		out.IntervalSeconds = nil
	}
	if in.TimeoutSeconds != nil {
		in, out := in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int64)
		**out = *in
	} else {
		out.TimeoutSeconds = nil
	}
	if err := intstr.DeepCopy_intstr_IntOrString(in.MaxUnavailable, &out.MaxUnavailable, c); err != nil {
		return err
	}
	if err := intstr.DeepCopy_intstr_IntOrString(in.MaxSurge, &out.MaxSurge, c); err != nil {
		return err
	}
	if in.UpdatePercent != nil {
		in, out := in.UpdatePercent, &out.UpdatePercent
		*out = new(int)
		**out = *in
	} else {
		out.UpdatePercent = nil
	}
	if in.Pre != nil {
		in, out := in.Pre, &out.Pre
		*out = new(LifecycleHook)
		if err := DeepCopy_api_LifecycleHook(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Post != nil {
		in, out := in.Post, &out.Post
		*out = new(LifecycleHook)
		if err := DeepCopy_api_LifecycleHook(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}

func DeepCopy_api_TagImageHook(in TagImageHook, out *TagImageHook, c *conversion.Cloner) error {
	out.ContainerName = in.ContainerName
	if err := api.DeepCopy_api_ObjectReference(in.To, &out.To, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_api_TemplateImage(in TemplateImage, out *TemplateImage, c *conversion.Cloner) error {
	out.Image = in.Image
	if in.Ref != nil {
		in, out := in.Ref, &out.Ref
		*out = new(image_api.DockerImageReference)
		if err := image_api.DeepCopy_api_DockerImageReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Ref = nil
	}
	if in.From != nil {
		in, out := in.From, &out.From
		*out = new(api.ObjectReference)
		if err := api.DeepCopy_api_ObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.From = nil
	}
	if in.Container != nil {
		in, out := in.Container, &out.Container
		*out = new(api.Container)
		if err := api.DeepCopy_api_Container(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Container = nil
	}
	return nil
}
