package authcmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aporeto.io/a3s/pkgs/authlib"
	"go.aporeto.io/a3s/pkgs/permissions"
	"go.aporeto.io/manipulate/manipcli"
)

func makeLDAPCmd(mmaker manipcli.ManipulatorMaker, restrictions *permissions.Restrictions) *cobra.Command {

	cmd := &cobra.Command{
		Use:              "ldap",
		Short:            "Use a configured LDAP authentication source.",
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			flags := cmd.Flags()
			fSourceName, _ := flags.GetString("source-name")
			fSourceNamespace, _ := flags.GetString("source-namespace")
			fAudience := viper.GetStringSlice("audience")
			fUser := viper.GetString("user")
			fPass := viper.GetString("pass")
			fCloak := viper.GetStringSlice("cloak")
			fQRCode := viper.GetBool("qrcode")
			fValidity := viper.GetDuration("validity")

			if fSourceNamespace == "" {
				fSourceNamespace = viper.GetString("namespace")
			}

			m, err := mmaker()
			if err != nil {
				return err
			}

			client := authlib.NewClient(m)
			t, err := client.AuthFromLDAP(
				context.Background(),
				fUser,
				fPass,
				fSourceNamespace,
				fSourceName,
				authlib.OptAudience(fAudience...),
				authlib.OptCloak(fCloak...),
				authlib.OptRestrictions(*restrictions),
				authlib.OptValidity(fValidity),
			)
			if err != nil {
				return err
			}

			printToken(t, fQRCode)

			return nil
		},
	}

	cmd.Flags().String("user", "", "The LDAP username to use.")
	cmd.Flags().String("pass", "", "The password associateds to the user.")
	_ = cobra.MarkFlagRequired(cmd.Flags(), "user")
	_ = cobra.MarkFlagRequired(cmd.Flags(), "pass")

	return cmd

}