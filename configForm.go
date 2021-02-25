package main

import (
	"net"
	"strconv"
	"strings"

	"github.com/andlabs/ui"
)

func makeConfigForm(window *ui.Window) (*ui.Form, *Config) {
	cfg := &Config{}

	configForm := ui.NewForm()
	configForm.SetPadded(true)

	ifacesBox := ui.NewCombobox()
	ifaceNames, err := GetAllIfaceNames(true)
	if err != nil {
		panic(err)
	}
	for _, ifaceName := range ifaceNames {
		ifacesBox.Append(ifaceName)
	}
	ifacesBox.SetSelected(0)
	cfg.Iface, _ = net.InterfaceByName(ifaceNames[ifacesBox.Selected()])

	IP, _ := getIPv4ByIface(cfg.Iface)
	myIPlabel := ui.NewLabel(IP.String())

	MAC := cfg.Iface.HardwareAddr
	myMAClabel := ui.NewLabel(MAC.String())
	//扫描方式
	methodBox := ui.NewCombobox()
	for _, v := range methods {
		methodBox.Append(string(v))
	}
	methodBox.SetSelected(0)
	cfg.ScanMethod = methods[0]
	methodBox.OnSelected(func(c *ui.Combobox) {
		index := c.Selected()
		cfg.ScanMethod = methods[index]
	})
	//隐蔽程度
	hideLevelBox := ui.NewCombobox()
	for _, v := range hideLevels {
		hideLevelBox.Append(string(v))
	}
	hideLevelBox.SetSelected(0)
	cfg.HideLevel = hideLevels[0]
	hideLevelBox.OnSelected(func(c *ui.Combobox) {
		index := c.Selected()
		cfg.HideLevel = hideLevels[index]
	})

	ifacesBox.OnSelected(func(c *ui.Combobox) {
		cfg.Iface, _ = net.InterfaceByIndex(ifacesBox.Selected())
		IP, err = getIPv4ByIface(cfg.Iface)
		if err != nil {
			ui.MsgBox(window, "Alert", "This interface don't have IP")
			return
		}
		myIPlabel.SetText(IP.String())
		myMAClabel.SetText(cfg.Iface.HardwareAddr.String())
	})

	rangeDiv := ui.NewHorizontalBox()

	nums := []*ui.Spinbox{}
	IPSlice := strings.Split(IP.String(), ".")
	count := 0
	for _, ipslice := range IPSlice {
		if count >= 3 {
			break
		}
		s := ui.NewSpinbox(0, 255)
		v, _ := strconv.Atoi(ipslice)
		s.SetValue(v)
		nums = append(nums, s)
		rangeDiv.Append(s, true)
		count++
	}

	cfg.MinBox = ui.NewSpinbox(0, 255)
	cfg.MaxBox = ui.NewSpinbox(0, 255)
	cfg.MinBox.SetValue(0)
	cfg.MaxBox.SetValue(255)
	rangeDiv.Append(cfg.MinBox, true)
	rangeDiv.Append(ui.NewLabel("~"), false)
	rangeDiv.Append(cfg.MaxBox, true)

	configForm.Append("网卡名称", ifacesBox, false)
	configForm.Append("我的IP", myIPlabel, false)
	configForm.Append("我的MAC地址", myMAClabel, false)
	configForm.Append("扫描范围", rangeDiv, false)
	configForm.Append("扫描方式", methodBox, false)
	configForm.Append("隐蔽程度", hideLevelBox, false)

	return configForm, cfg
}
