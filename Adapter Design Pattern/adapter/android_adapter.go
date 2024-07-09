package adapter

import "AdapterPattern/services"

type AndroidAdapter struct {
	Android *services.Android
}

func (ad *AndroidAdapter) ChargeAppleMobile() {
	ad.Android.ChargeAndroidMobile()
}
