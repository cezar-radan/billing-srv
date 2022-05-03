/*
Copyright 2022.
This file uses pdfcpu as PDF processor (https://pdfcpu.io/).

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
	http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"github.com/ledgertech/billing-srv/util/assets"
	"github.com/ledgertech/billing-srv/util/clog"
	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
)

//func check/install user fonts (embedded fonts)
func handleEmbededFonts() error {
	pdfapi.LoadConfiguration()

	fontsInstalled, err := pdfapi.ListFonts()
	if err != nil {
		clog.Errorf("cannot list fonts %+v\n", err)
		return err
	}
	clog.Info("-----------------------")
	clog.Infof("installed fonts before:%+v", fontsInstalled)
	clog.Info("-----------------------")

	fontsEmbedded, err := assets.UserFonts("util/assets/fonts")
	if err != nil {
		clog.Errorf("cannot load embeded fonts: %+v\n", err)
		return err
	}
	clog.Info("-----------------------")
	clog.Infof("embedded fonts:%+v", fontsEmbedded)
	clog.Info("-----------------------")

	// install fonts
	if len(fontsEmbedded) > 0 {
		err = pdfapi.InstallFonts(fontsEmbedded)
		if err != nil {
			clog.Errorf("cannot install fonts: %+v\n", err)
			return err
		}
	}

	fontsInstalled, err = pdfapi.ListFonts()
	if err != nil {
		clog.Errorf("cannot list fonts %+v\n", err)
		return err
	}
	clog.Info("-----------------------")
	clog.Infof("installed fonts after:%+v", fontsInstalled)
	clog.Info("-----------------------")

	return nil
}
