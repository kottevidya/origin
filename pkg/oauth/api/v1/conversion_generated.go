// +build !ignore_autogenerated

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	oauth_api "github.com/openshift/origin/pkg/oauth/api"
	api "k8s.io/kubernetes/pkg/api"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedConversionFuncs(
		Convert_v1_OAuthAccessToken_To_api_OAuthAccessToken,
		Convert_api_OAuthAccessToken_To_v1_OAuthAccessToken,
		Convert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList,
		Convert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList,
		Convert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken,
		Convert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken,
		Convert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList,
		Convert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList,
		Convert_v1_OAuthClient_To_api_OAuthClient,
		Convert_api_OAuthClient_To_v1_OAuthClient,
		Convert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization,
		Convert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization,
		Convert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList,
		Convert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList,
		Convert_v1_OAuthClientList_To_api_OAuthClientList,
		Convert_api_OAuthClientList_To_v1_OAuthClientList,
	); err != nil {
		// if one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

func autoConvert_v1_OAuthAccessToken_To_api_OAuthAccessToken(in *OAuthAccessToken, out *oauth_api.OAuthAccessToken, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*OAuthAccessToken))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.AuthorizeToken = in.AuthorizeToken
	out.RefreshToken = in.RefreshToken
	return nil
}

func Convert_v1_OAuthAccessToken_To_api_OAuthAccessToken(in *OAuthAccessToken, out *oauth_api.OAuthAccessToken, s conversion.Scope) error {
	return autoConvert_v1_OAuthAccessToken_To_api_OAuthAccessToken(in, out, s)
}

func autoConvert_api_OAuthAccessToken_To_v1_OAuthAccessToken(in *oauth_api.OAuthAccessToken, out *OAuthAccessToken, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauth_api.OAuthAccessToken))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.AuthorizeToken = in.AuthorizeToken
	out.RefreshToken = in.RefreshToken
	return nil
}

func Convert_api_OAuthAccessToken_To_v1_OAuthAccessToken(in *oauth_api.OAuthAccessToken, out *OAuthAccessToken, s conversion.Scope) error {
	return autoConvert_api_OAuthAccessToken_To_v1_OAuthAccessToken(in, out, s)
}

func autoConvert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList(in *OAuthAccessTokenList, out *oauth_api.OAuthAccessTokenList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*OAuthAccessTokenList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]oauth_api.OAuthAccessToken, len(*in))
		for i := range *in {
			if err := Convert_v1_OAuthAccessToken_To_api_OAuthAccessToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList(in *OAuthAccessTokenList, out *oauth_api.OAuthAccessTokenList, s conversion.Scope) error {
	return autoConvert_v1_OAuthAccessTokenList_To_api_OAuthAccessTokenList(in, out, s)
}

func autoConvert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in *oauth_api.OAuthAccessTokenList, out *OAuthAccessTokenList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauth_api.OAuthAccessTokenList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthAccessToken, len(*in))
		for i := range *in {
			if err := Convert_api_OAuthAccessToken_To_v1_OAuthAccessToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in *oauth_api.OAuthAccessTokenList, out *OAuthAccessTokenList, s conversion.Scope) error {
	return autoConvert_api_OAuthAccessTokenList_To_v1_OAuthAccessTokenList(in, out, s)
}

func autoConvert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(in *OAuthAuthorizeToken, out *oauth_api.OAuthAuthorizeToken, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*OAuthAuthorizeToken))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.State = in.State
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	return nil
}

func Convert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(in *OAuthAuthorizeToken, out *oauth_api.OAuthAuthorizeToken, s conversion.Scope) error {
	return autoConvert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(in, out, s)
}

func autoConvert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in *oauth_api.OAuthAuthorizeToken, out *OAuthAuthorizeToken, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauth_api.OAuthAuthorizeToken))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.State = in.State
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	return nil
}

func Convert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in *oauth_api.OAuthAuthorizeToken, out *OAuthAuthorizeToken, s conversion.Scope) error {
	return autoConvert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(in, out, s)
}

func autoConvert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList(in *OAuthAuthorizeTokenList, out *oauth_api.OAuthAuthorizeTokenList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*OAuthAuthorizeTokenList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]oauth_api.OAuthAuthorizeToken, len(*in))
		for i := range *in {
			if err := Convert_v1_OAuthAuthorizeToken_To_api_OAuthAuthorizeToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList(in *OAuthAuthorizeTokenList, out *oauth_api.OAuthAuthorizeTokenList, s conversion.Scope) error {
	return autoConvert_v1_OAuthAuthorizeTokenList_To_api_OAuthAuthorizeTokenList(in, out, s)
}

func autoConvert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in *oauth_api.OAuthAuthorizeTokenList, out *OAuthAuthorizeTokenList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauth_api.OAuthAuthorizeTokenList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthAuthorizeToken, len(*in))
		for i := range *in {
			if err := Convert_api_OAuthAuthorizeToken_To_v1_OAuthAuthorizeToken(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in *oauth_api.OAuthAuthorizeTokenList, out *OAuthAuthorizeTokenList, s conversion.Scope) error {
	return autoConvert_api_OAuthAuthorizeTokenList_To_v1_OAuthAuthorizeTokenList(in, out, s)
}

func autoConvert_v1_OAuthClient_To_api_OAuthClient(in *OAuthClient, out *oauth_api.OAuthClient, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*OAuthClient))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.Secret = in.Secret
	out.RespondWithChallenges = in.RespondWithChallenges
	if in.RedirectURIs != nil {
		in, out := &in.RedirectURIs, &out.RedirectURIs
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.RedirectURIs = nil
	}
	return nil
}

func Convert_v1_OAuthClient_To_api_OAuthClient(in *OAuthClient, out *oauth_api.OAuthClient, s conversion.Scope) error {
	return autoConvert_v1_OAuthClient_To_api_OAuthClient(in, out, s)
}

func autoConvert_api_OAuthClient_To_v1_OAuthClient(in *oauth_api.OAuthClient, out *OAuthClient, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauth_api.OAuthClient))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.Secret = in.Secret
	out.RespondWithChallenges = in.RespondWithChallenges
	if in.RedirectURIs != nil {
		in, out := &in.RedirectURIs, &out.RedirectURIs
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.RedirectURIs = nil
	}
	return nil
}

func Convert_api_OAuthClient_To_v1_OAuthClient(in *oauth_api.OAuthClient, out *OAuthClient, s conversion.Scope) error {
	return autoConvert_api_OAuthClient_To_v1_OAuthClient(in, out, s)
}

func autoConvert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(in *OAuthClientAuthorization, out *oauth_api.OAuthClientAuthorization, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*OAuthClientAuthorization))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Scopes = nil
	}
	return nil
}

func Convert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(in *OAuthClientAuthorization, out *oauth_api.OAuthClientAuthorization, s conversion.Scope) error {
	return autoConvert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(in, out, s)
}

func autoConvert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in *oauth_api.OAuthClientAuthorization, out *OAuthClientAuthorization, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauth_api.OAuthClientAuthorization))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.ClientName = in.ClientName
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Scopes = nil
	}
	return nil
}

func Convert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in *oauth_api.OAuthClientAuthorization, out *OAuthClientAuthorization, s conversion.Scope) error {
	return autoConvert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(in, out, s)
}

func autoConvert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList(in *OAuthClientAuthorizationList, out *oauth_api.OAuthClientAuthorizationList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*OAuthClientAuthorizationList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]oauth_api.OAuthClientAuthorization, len(*in))
		for i := range *in {
			if err := Convert_v1_OAuthClientAuthorization_To_api_OAuthClientAuthorization(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList(in *OAuthClientAuthorizationList, out *oauth_api.OAuthClientAuthorizationList, s conversion.Scope) error {
	return autoConvert_v1_OAuthClientAuthorizationList_To_api_OAuthClientAuthorizationList(in, out, s)
}

func autoConvert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in *oauth_api.OAuthClientAuthorizationList, out *OAuthClientAuthorizationList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauth_api.OAuthClientAuthorizationList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthClientAuthorization, len(*in))
		for i := range *in {
			if err := Convert_api_OAuthClientAuthorization_To_v1_OAuthClientAuthorization(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in *oauth_api.OAuthClientAuthorizationList, out *OAuthClientAuthorizationList, s conversion.Scope) error {
	return autoConvert_api_OAuthClientAuthorizationList_To_v1_OAuthClientAuthorizationList(in, out, s)
}

func autoConvert_v1_OAuthClientList_To_api_OAuthClientList(in *OAuthClientList, out *oauth_api.OAuthClientList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*OAuthClientList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]oauth_api.OAuthClient, len(*in))
		for i := range *in {
			if err := Convert_v1_OAuthClient_To_api_OAuthClient(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_OAuthClientList_To_api_OAuthClientList(in *OAuthClientList, out *oauth_api.OAuthClientList, s conversion.Scope) error {
	return autoConvert_v1_OAuthClientList_To_api_OAuthClientList(in, out, s)
}

func autoConvert_api_OAuthClientList_To_v1_OAuthClientList(in *oauth_api.OAuthClientList, out *OAuthClientList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*oauth_api.OAuthClientList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OAuthClient, len(*in))
		for i := range *in {
			if err := Convert_api_OAuthClient_To_v1_OAuthClient(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_OAuthClientList_To_v1_OAuthClientList(in *oauth_api.OAuthClientList, out *OAuthClientList, s conversion.Scope) error {
	return autoConvert_api_OAuthClientList_To_v1_OAuthClientList(in, out, s)
}
