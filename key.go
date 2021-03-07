package samsung

type Key string

const (
	// Power
	KEY_POWEROFF Key = "KEY_POWEROFF"
	KEY_POWERON  Key = "KEY_POWERON"
	KEY_POWER    Key = "KEY_POWER"

	// Input
	KEY_SOURCE     Key = "KEY_SOURCE"
	KEY_COMPONENT1 Key = "KEY_COMPONENT1"
	KEY_COMPONENT2 Key = "KEY_COMPONENT2"
	KEY_AV1        Key = "KEY_AV1"
	KEY_AV2        Key = "KEY_AV2"
	KEY_AV3        Key = "KEY_AV3"
	KEY_SVIDEO1    Key = "KEY_SVIDEO1"
	KEY_SVIDEO2    Key = "KEY_SVIDEO2"
	KEY_SVIDEO3    Key = "KEY_SVIDEO3"
	KEY_HDMI       Key = "KEY_HDMI"
	KEY_HDMI1      Key = "KEY_HDMI1"
	KEY_HDMI2      Key = "KEY_HDMI2"
	KEY_HDMI3      Key = "KEY_HDMI3"
	KEY_HDMI4      Key = "KEY_HDMI4"
	KEY_FM_RADIO   Key = "KEY_FM_RADIO"
	KEY_DVI        Key = "KEY_DVI"
	KEY_DVR        Key = "KEY_DVR"
	KEY_TV         Key = "KEY_TV"
	KEY_ANTENA     Key = "KEY_ANTENA"
	KEY_DTV        Key = "KEY_DTV"

	// Number
	KEY_1 Key = "KEY_1"
	KEY_2 Key = "KEY_2"
	KEY_3 Key = "KEY_3"
	KEY_4 Key = "KEY_4"
	KEY_5 Key = "KEY_5"
	KEY_6 Key = "KEY_6"
	KEY_7 Key = "KEY_7"
	KEY_8 Key = "KEY_8"
	KEY_9 Key = "KEY_9"
	KEY_0 Key = "KEY_0"

	// Misc
	KEY_PANNEL_CHDOWN Key = "KEY_PANNEL_CHDOWN"
	KEY_ANYNET        Key = "KEY_ANYNET"
	KEY_ESAVING       Key = "KEY_ESAVING"
	KEY_SLEEP         Key = "KEY_SLEEP"
	KEY_DTV_SIGNAL    Key = "KEY_DTV_SIGNAL"

	// Channel
	KEY_CHUP          Key = "KEY_CHUP"
	KEY_CHDOWN        Key = "KEY_CHDOWN"
	KEY_PRECH         Key = "KEY_PRECH"
	KEY_FAVCH         Key = "KEY_FAVCH"
	KEY_CH_LIST       Key = "KEY_CH_LIST"
	KEY_AUTO_PROGRAM  Key = "KEY_AUTO_PROGRAM"
	KEY_MAGIC_CHANNEL Key = "KEY_MAGIC_CHANNEL"

	// Volume
	KEY_VOLUP   Key = "KEY_VOLUP"
	KEY_VOLDOWN Key = "KEY_VOLDOWN"
	KEY_MUTE    Key = "KEY_MUTE"

	// Direction
	KEY_UP     Key = "KEY_UP"
	KEY_DOWN   Key = "KEY_DOWN"
	KEY_LEFT   Key = "KEY_LEFT"
	KEY_RIGHT  Key = "KEY_RIGHT"
	KEY_RETURN Key = "KEY_RETURN"
	KEY_ENTER  Key = "KEY_ENTER"

	// Media
	KEY_REWIND         Key = "KEY_REWIND"
	KEY_STOP           Key = "KEY_STOP"
	KEY_PLAY           Key = "KEY_PLAY"
	KEY_FF             Key = "KEY_FF"
	KEY_REC            Key = "KEY_REC"
	KEY_PAUSE          Key = "KEY_PAUSE"
	KEY_LIVE           Key = "KEY_LIVE"
	KEY_QUICK_REPLAY   Key = "KEY_QUICK_REPLAY"
	KEY_STILL_PICTURE  Key = "KEY_STILL_PICTURE"
	KEY_INSTANT_REPLAY Key = "KEY_INSTANT_REPLAY"

	// Picture in picture
	KEY_PIP_ONOFF                  Key = "KEY_PIP_ONOFF"
	KEY_PIP_SWAP                   Key = "KEY_PIP_SWAP"
	KEY_PIP_SIZE                   Key = "KEY_PIP_SIZE"
	KEY_PIP_CHUP                   Key = "KEY_PIP_CHUP"
	KEY_PIP_CHDOWN                 Key = "KEY_PIP_CHDOWN"
	KEY_AUTO_ARC_PIP_SMALL         Key = "KEY_AUTO_ARC_PIP_SMALL"
	KEY_AUTO_ARC_PIP_WIDE          Key = "KEY_AUTO_ARC_PIP_WIDE"
	KEY_AUTO_ARC_PIP_RIGHT_BOTTOM  Key = "KEY_AUTO_ARC_PIP_RIGHT_BOTTOM"
	KEY_AUTO_ARC_PIP_SOURCE_CHANGE Key = "KEY_AUTO_ARC_PIP_SOURCE_CHANGE"
	KEY_PIP_SCAN                   Key = "KEY_PIP_SCAN"

	// Mode
	KEY_VCR_MODE  Key = "KEY_VCR_MODE"
	KEY_CATV_MODE Key = "KEY_CATV_MODE"
	KEY_DSS_MODE  Key = "KEY_DSS_MODE"
	KEY_TV_MODE   Key = "KEY_TV_MODE"
	KEY_DVD_MODE  Key = "KEY_DVD_MODE"
	KEY_STB_MODE  Key = "KEY_STB_MODE"
	KEY_PCMODE    Key = "KEY_PCMODE"

	// Color
	KEY_GREEN  Key = "KEY_GREEN"
	KEY_YELLOW Key = "KEY_YELLOW"
	KEY_CYAN   Key = "KEY_CYAN"
	KEY_RED    Key = "KEY_RED"

	// Teletext
	KEY_TTX_MIX     Key = "KEY_TTX_MIX"
	KEY_TTX_SUBFACE Key = "KEY_TTX_SUBFACE"

	// Aspect ratio
	KEY_ASPECT       Key = "KEY_ASPECT"
	KEY_PICTURE_SIZE Key = "KEY_PICTURE_SIZE"
	KEY_4_3          Key = "KEY_4_3"
	KEY_16_9         Key = "KEY_16_9"
	KEY_EXT14        Key = "KEY_EXT14"
	KEY_EXT15        Key = "KEY_EXT15"

	// Picture mode
	KEY_PMODE    Key = "KEY_PMODE"
	KEY_PANORAMA Key = "KEY_PANORAMA"
	KEY_DYNAMIC  Key = "KEY_DYNAMIC"
	KEY_STANDARD Key = "KEY_STANDARD"
	KEY_MOVIE1   Key = "KEY_MOVIE1"
	KEY_GAME     Key = "KEY_GAME"
	KEY_CUSTOM   Key = "KEY_CUSTOM"
	KEY_EXT9     Key = "KEY_EXT9"
	KEY_EXT10    Key = "KEY_EXT10"

	// Menu
	KEY_MENU      Key = "KEY_MENU"
	KEY_TOPMENU   Key = "KEY_TOPMENU"
	KEY_TOOLS     Key = "KEY_TOOLS"
	KEY_HOME      Key = "KEY_HOME"
	KEY_CONTENTS  Key = "KEY_CONTENTS"
	KEY_GUIDE     Key = "KEY_GUIDE"
	KEY_DISC_MENU Key = "KEY_DISC_MENU"
	KEY_DVR_MENU  Key = "KEY_DVR_MENU"
	KEY_HELP      Key = "KEY_HELP"

	// OSD
	KEY_INFO              Key = "KEY_INFO"
	KEY_CAPTION           Key = "KEY_CAPTION"
	KEY_CLOCK_DISPLAY     Key = "KEY_CLOCK_DISPLAY"
	KEY_SETUP_CLOCK_TIMER Key = "KEY_SETUP_CLOCK_TIMER"
	KEY_SUB_TITLE         Key = "KEY_SUB_TITLE"

	// Zoom
	KEY_ZOOM_MOVE Key = "KEY_ZOOM_MOVE"
	KEY_ZOOM_IN   Key = "KEY_ZOOM_IN"
	KEY_ZOOM_OUT  Key = "KEY_ZOOM_OUT"
	KEY_ZOOM1     Key = "KEY_ZOOM1"
	KEY_ZOOM2     Key = "KEY_ZOOM2"

	// Other
	KEY_WHEEL_LEFT            Key = "KEY_WHEEL_LEFT"
	KEY_WHEEL_RIGHT           Key = "KEY_WHEEL_RIGHT"
	KEY_ADDDEL                Key = "KEY_ADDDEL"
	KEY_PLUS100               Key = "KEY_PLUS100"
	KEY_AD                    Key = "KEY_AD"
	KEY_LINK                  Key = "KEY_LINK"
	KEY_TURBO                 Key = "KEY_TURBO"
	KEY_CONVERGENCE           Key = "KEY_CONVERGENCE"
	KEY_DEVICE_CONNECT        Key = "KEY_DEVICE_CONNECT"
	KEY_11                    Key = "KEY_11"
	KEY_12                    Key = "KEY_12"
	KEY_FACTORY               Key = "KEY_FACTORY"
	KEY_3SPEED                Key = "KEY_3SPEED"
	KEY_RSURF                 Key = "KEY_RSURF"
	KEY_FF_                   Key = "KEY_FF_"
	KEY_REWIND_               Key = "KEY_REWIND_"
	KEY_ANGLE                 Key = "KEY_ANGLE"
	KEY_RESERVED1             Key = "KEY_RESERVED1"
	KEY_PROGRAM               Key = "KEY_PROGRAM"
	KEY_BOOKMARK              Key = "KEY_BOOKMARK"
	KEY_PRINT                 Key = "KEY_PRINT"
	KEY_CLEAR                 Key = "KEY_CLEAR"
	KEY_VCHIP                 Key = "KEY_VCHIP"
	KEY_REPEAT                Key = "KEY_REPEAT"
	KEY_DOOR                  Key = "KEY_DOOR"
	KEY_OPEN                  Key = "KEY_OPEN"
	KEY_DMA                   Key = "KEY_DMA"
	KEY_MTS                   Key = "KEY_MTS"
	KEY_DNIe                  Key = "KEY_DNIe"
	KEY_SRS                   Key = "KEY_SRS"
	KEY_CONVERT_AUDIO_MAINSUB Key = "KEY_CONVERT_AUDIO_MAINSUB"
	KEY_MDC                   Key = "KEY_MDC"
	KEY_SEFFECT               Key = "KEY_SEFFECT"
	KEY_PERPECT_FOCUS         Key = "KEY_PERPECT_FOCUS"
	KEY_CALLER_ID             Key = "KEY_CALLER_ID"
	KEY_SCALE                 Key = "KEY_SCALE"
	KEY_MAGIC_BRIGHT          Key = "KEY_MAGIC_BRIGHT"
	KEY_W_LINK                Key = "KEY_W_LINK"
	KEY_DTV_LINK              Key = "KEY_DTV_LINK"
	KEY_APP_LIST              Key = "KEY_APP_LIST"
	KEY_BACK_MHP              Key = "KEY_BACK_MHP"
	KEY_ALT_MHP               Key = "KEY_ALT_MHP"
	KEY_DNSe                  Key = "KEY_DNSe"
	KEY_RSS                   Key = "KEY_RSS"
	KEY_ENTERTAINMENT         Key = "KEY_ENTERTAINMENT"
	KEY_ID_INPUT              Key = "KEY_ID_INPUT"
	KEY_ID_SETUP              Key = "KEY_ID_SETUP"
	KEY_ANYVIEW               Key = "KEY_ANYVIEW"
	KEY_MS                    Key = "KEY_MS"
	KEY_MORE                  Key = "KEY_MORE"
	KEY_MIC                   Key = "KEY_MIC"
	KEY_NINE_SEPERATE         Key = "KEY_NINE_SEPERATE"
	KEY_AUTO_FORMAT           Key = "KEY_AUTO_FORMAT"
	KEY_DNET                  Key = "KEY_DNET"

	// Auto arc
	KEY_AUTO_ARC_C_FORCE_AGING     Key = "KEY_AUTO_ARC_C_FORCE_AGING"
	KEY_AUTO_ARC_CAPTION_ENG       Key = "KEY_AUTO_ARC_CAPTION_ENG"
	KEY_AUTO_ARC_USBJACK_INSPECT   Key = "KEY_AUTO_ARC_USBJACK_INSPECT"
	KEY_AUTO_ARC_RESET             Key = "KEY_AUTO_ARC_RESET"
	KEY_AUTO_ARC_LNA_ON            Key = "KEY_AUTO_ARC_LNA_ON"
	KEY_AUTO_ARC_LNA_OFF           Key = "KEY_AUTO_ARC_LNA_OFF"
	KEY_AUTO_ARC_ANYNET_MODE_OK    Key = "KEY_AUTO_ARC_ANYNET_MODE_OK"
	KEY_AUTO_ARC_ANYNET_AUTO_START Key = "KEY_AUTO_ARC_ANYNET_AUTO_START"
	KEY_AUTO_ARC_CAPTION_ON        Key = "KEY_AUTO_ARC_CAPTION_ON"
	KEY_AUTO_ARC_CAPTION_OFF       Key = "KEY_AUTO_ARC_CAPTION_OFF"
	KEY_AUTO_ARC_PIP_DOUBLE        Key = "KEY_AUTO_ARC_PIP_DOUBLE"
	KEY_AUTO_ARC_PIP_LARGE         Key = "KEY_AUTO_ARC_PIP_LARGE"
	KEY_AUTO_ARC_PIP_LEFT_TOP      Key = "KEY_AUTO_ARC_PIP_LEFT_TOP"
	KEY_AUTO_ARC_PIP_RIGHT_TOP     Key = "KEY_AUTO_ARC_PIP_RIGHT_TOP"
	KEY_AUTO_ARC_PIP_LEFT_BOTTOM   Key = "KEY_AUTO_ARC_PIP_LEFT_BOTTOM"
	KEY_AUTO_ARC_PIP_CH_CHANGE     Key = "KEY_AUTO_ARC_PIP_CH_CHANGE"
	KEY_AUTO_ARC_AUTOCOLOR_SUCCESS Key = "KEY_AUTO_ARC_AUTOCOLOR_SUCCESS"
	KEY_AUTO_ARC_AUTOCOLOR_FAIL    Key = "KEY_AUTO_ARC_AUTOCOLOR_FAIL"
	KEY_AUTO_ARC_JACK_IDENT        Key = "KEY_AUTO_ARC_JACK_IDENT"
	KEY_AUTO_ARC_CAPTION_KOR       Key = "KEY_AUTO_ARC_CAPTION_KOR"
	KEY_AUTO_ARC_ANTENNA_AIR       Key = "KEY_AUTO_ARC_ANTENNA_AIR"
	KEY_AUTO_ARC_ANTENNA_CABLE     Key = "KEY_AUTO_ARC_ANTENNA_CABLE"
	KEY_AUTO_ARC_ANTENNA_SATELLITE Key = "KEY_AUTO_ARC_ANTENNA_SATELLITE"

	// Panel
	KEY_PANNEL_POWER  Key = "KEY_PANNEL_POWER"
	KEY_PANNEL_CHUP   Key = "KEY_PANNEL_CHUP"
	KEY_PANNEL_VOLUP  Key = "KEY_PANNEL_VOLUP"
	KEY_PANNEL_VOLDOW Key = "KEY_PANNEL_VOLDOW"
	KEY_PANNEL_ENTER  Key = "KEY_PANNEL_ENTER"
	KEY_PANNEL_MENU   Key = "KEY_PANNEL_MENU"
	KEY_PANNEL_SOURCE Key = "KEY_PANNEL_SOURCE"

	// Extended
	KEY_EXT1  Key = "KEY_EXT1"
	KEY_EXT2  Key = "KEY_EXT2"
	KEY_EXT3  Key = "KEY_EXT3"
	KEY_EXT4  Key = "KEY_EXT4"
	KEY_EXT5  Key = "KEY_EXT5"
	KEY_EXT6  Key = "KEY_EXT6"
	KEY_EXT7  Key = "KEY_EXT7"
	KEY_EXT8  Key = "KEY_EXT8"
	KEY_EXT11 Key = "KEY_EXT11"
	KEY_EXT12 Key = "KEY_EXT12"
	KEY_EXT13 Key = "KEY_EXT13"
	KEY_EXT16 Key = "KEY_EXT16"
	KEY_EXT17 Key = "KEY_EXT17"
	KEY_EXT18 Key = "KEY_EXT18"
	KEY_EXT19 Key = "KEY_EXT19"
	KEY_EXT20 Key = "KEY_EXT20"
	KEY_EXT21 Key = "KEY_EXT21"
	KEY_EXT22 Key = "KEY_EXT22"
	KEY_EXT23 Key = "KEY_EXT23"
	KEY_EXT24 Key = "KEY_EXT24"
	KEY_EXT25 Key = "KEY_EXT25"
	KEY_EXT26 Key = "KEY_EXT26"
	KEY_EXT27 Key = "KEY_EXT27"
	KEY_EXT28 Key = "KEY_EXT28"
	KEY_EXT29 Key = "KEY_EXT29"
	KEY_EXT30 Key = "KEY_EXT30"
	KEY_EXT31 Key = "KEY_EXT31"
	KEY_EXT32 Key = "KEY_EXT32"
	KEY_EXT33 Key = "KEY_EXT33"
	KEY_EXT34 Key = "KEY_EXT34"
	KEY_EXT35 Key = "KEY_EXT35"
	KEY_EXT36 Key = "KEY_EXT36"
	KEY_EXT37 Key = "KEY_EXT37"
	KEY_EXT38 Key = "KEY_EXT38"
	KEY_EXT39 Key = "KEY_EXT39"
	KEY_EXT40 Key = "KEY_EXT40"
	KEY_EXT41 Key = "KEY_EXT41"
)