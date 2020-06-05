// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package baremetalproviderspec

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServer) DeepCopyInto(out *APIServer) {
	*out = *in
	if in.AdditionalSANs != nil {
		in, out := &in.AdditionalSANs, &out.AdditionalSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServer.
func (in *APIServer) DeepCopy() *APIServer {
	if in == nil {
		return nil
	}
	out := new(APIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationWebhook) DeepCopyInto(out *AuthenticationWebhook) {
	*out = *in
	in.WebhookClient.DeepCopyInto(&out.WebhookClient)
	in.WebhookServer.DeepCopyInto(&out.WebhookServer)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationWebhook.
func (in *AuthenticationWebhook) DeepCopy() *AuthenticationWebhook {
	if in == nil {
		return nil
	}
	out := new(AuthenticationWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationWebhook) DeepCopyInto(out *AuthorizationWebhook) {
	*out = *in
	in.WebhookClient.DeepCopyInto(&out.WebhookClient)
	in.WebhookServer.DeepCopyInto(&out.WebhookServer)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationWebhook.
func (in *AuthorizationWebhook) DeepCopy() *AuthorizationWebhook {
	if in == nil {
		return nil
	}
	out := new(AuthorizationWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalClusterProviderSpec) DeepCopyInto(out *BareMetalClusterProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(AuthenticationWebhook)
		(*in).DeepCopyInto(*out)
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(AuthorizationWebhook)
		(*in).DeepCopyInto(*out)
	}
	in.APIServer.DeepCopyInto(&out.APIServer)
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalClusterProviderSpec.
func (in *BareMetalClusterProviderSpec) DeepCopy() *BareMetalClusterProviderSpec {
	if in == nil {
		return nil
	}
	out := new(BareMetalClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalClusterProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalMachineProviderSpec) DeepCopyInto(out *BareMetalMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Private = in.Private
	out.Public = in.Public
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalMachineProviderSpec.
func (in *BareMetalMachineProviderSpec) DeepCopy() *BareMetalMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndPoint) DeepCopyInto(out *EndPoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndPoint.
func (in *EndPoint) DeepCopy() *EndPoint {
	if in == nil {
		return nil
	}
	out := new(EndPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookClient) DeepCopyInto(out *WebhookClient) {
	*out = *in
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CertificateData != nil {
		in, out := &in.CertificateData, &out.CertificateData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookClient.
func (in *WebhookClient) DeepCopy() *WebhookClient {
	if in == nil {
		return nil
	}
	out := new(WebhookClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookServer) DeepCopyInto(out *WebhookServer) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookServer.
func (in *WebhookServer) DeepCopy() *WebhookServer {
	if in == nil {
		return nil
	}
	out := new(WebhookServer)
	in.DeepCopyInto(out)
	return out
}