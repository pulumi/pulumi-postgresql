## HEAD (Unreleased)
_(none)_

---

## 3.1.0 (2021-05-27)
* Upgrade to v1.13.0 Postgresql Terraform Provider

## 3.0.1 (2021-04-30)
* Upgrade to v1.12.1 Postgresql Terraform Provider

## 3.0.0 (2021-04-19)
* Depend on Pulumi 3.0, which includes improvements to Python resource arguments and key translation, Go SDK performance,
  Node SDK performance, general availability of Automation API, and more.

## 2.10.0 (2021-04-12)
* Upgrade to pulumi-terraform-bridge v2.23.0

## 2.9.0 (2021-04-05)
* Upgrade to v1.12.0 Postgresql Terraform Provider

## 2.8.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 2.8.0 (2021-03-15)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 2.7.2 (2021-02-18)
* Upgrade to v1.11.2 of the Postgresql Terraform Provider

## 2.7.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 2.7.0 (2021-02-03)
* Upgrade to v1.11.1 of the Postgresql Terraform Provider

## 2.6.0 (2021-01-29)
* Upgrading pulumi-terraform-bridge to v2.18.1

## 2.5.3 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 2.5.2 (2020-11-23)
* Upgrade to pulumi-terraform-bridge v2.13.2  
  * This adds support for import specific examples in documentation

## 2.5.1 (2020-11-09)
* Upgrade to pulumi-terraform-bridge v2.12.1

## 2.5.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.4.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 2.3.0 (2020-07-30)
* Upgrade to v1.7.1 of the Postgresql Terraform Provider

## 2.2.2 (2020-06-12)
* Switch to GitHub actions for build

## 2.2.1 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.2.0 (2020-05-27)
* Upgrade to pulumi-terraform-bridge v2.3.3
* Upgrade to v1.6.0 of the postgresql Terraform Provider

## 2.1.1 (2020-05-13)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.1.0 (2020-04-28)
* Upgrade to pulumi-terraform-bridge v2.1.1

## 2.0.0 (2020-04-17)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.7.0 (2020-04-14)
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge v1.8.4
* Refactor layout to support Go modules

## 1.6.0 (2020-03-23)
* Rename `defaultPrivileg` to `defaultPrivileges`
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## v1.5.0 (2020-02-25)
* Upgrade to v1.5.0 of the Postgresql Terraform Provider

## 1.4.0 (2020-01-29)
* Upgrade to pulumi-terraform-bridge v1.6.4

## 1.3.0 (2019-12-27)
* Namespace names in .NET SDK are adjusted to PascalCase
([#28](https://github.com/pulumi/pulumi-postgresql/pull/28)).
* Upgrade to pulumi-terraform-bridge v1.5.2
* Upgrade to v1.4.0 of the Postgresql Terraform Provider

## 1.2.0 (2019-12-04)
* Upgrade to support go 1.13.x
* Upgrade to pulumi-terraform-bridge v1.4.3

## 1.1.0 (2019-11-15)
* Add support for DotNet SDK Generation
* Upgrade to v1.3.0 of the Postgresql Terraform Provider

## 1.0.0 (2019-10-04)
* Regenerate SDK based on tf2pulumi 0.6.0
* Upgrade to v1.2.0 of the Postgresql Terraform Providers

## 0.18.9 (2019-09-05)
* Upgrade to Pulumi v1.0.0

## 0.18.8 (2019-08-29)
* Upgrade pulumi-terraform to 3f206601e7

## 0.18.7 (2019-08-20)
* Depend on latest pulumi package

## 0.18.6 (2019-08-09)
* Upgrade to pulumi-terraform 9db2fc93cd

## 0.18.5 (2019-08-08)
* Update to pulumi-terraform@013b95b1c

## 0.18.4 (2019-07-09)
* Fix detailed diffs with nested computed values.

## 0.18.3 (2019-07-08)
* Communicate detailed information about the difference between a resource's desired and actual state during a Pulumi update

## 0.18.2 (2019-07-04)
* Upgrade to `v1.1.0` of the postgresql Terraform provider.

## 0.18.1 (2019-06-21)
* Update to pulumi-terraform@3635bed3a5 which stops maps containing `.` being treated as nested maps.

## 0.18.0 (2019-06-17)
* Initial release of `pulumi-postgresql`, based on `v0.4.0` of the postgresql Terraform provider.
