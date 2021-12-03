// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expression

import (
	"regexp"
)

var (
	ipv4      = regexp.MustCompile(`^\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}$`)
	cellphone = regexp.MustCompile(`^(\+86(-|\s)?)?1[3-9][0-9]{9}$`)
	telephone = regexp.MustCompile(`^(\d{3,4}-)?[1-9][0-9]{4,8}$`)
	email     = regexp.MustCompile(`^[a-zA-Z0-9]{1,20}(([._\-])?[a-zA-Z0-9]{1,20})*@[a-z0-9]+(\.?[a-z]+)+$`)
)

// IsIPv4 regular expression is used to check whether the IP address is IPV4.
func IsIPv4(str string) bool {
	return ipv4.MatchString(str)
}

// IsCellphoneOfChina regular expression is used to check whether the telephone number is china.
func IsCellphoneOfChina(str string) bool {
	return cellphone.MatchString(str)
}

// IsTelephoneOfChina regular expression is used to check whether the telephone number is china.
func IsTelephoneOfChina(str string) bool {
	return telephone.MatchString(str)
}

// IsEmail Use a regular expression to check whether the email address is valid.
func IsEmail(str string) bool {
	return email.MatchString(str)
}