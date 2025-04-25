declare interface AdminConfigurations {
    username: string;
    password: string;
    google_captcha_site_key: string | undefined;
    google_captcha_secret_key: string | undefined;
}


export type {
    AdminConfigurations,
}