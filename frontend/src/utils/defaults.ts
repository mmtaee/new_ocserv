import type {PanelSetupData} from "@/api";

const panelSetupDefault: PanelSetupData = {
    admin: {
        username: '',
        password: '',
    },
    config: {
        google_captcha_secret_key: "",
        google_captcha_site_key: ""
    },
    default_ocserv_group: {}
}


export {
    panelSetupDefault
}