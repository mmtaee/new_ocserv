type Validator = (value: string, t: (key: string) => string) => true | string

const ipOrRangeRule: Validator = (v, t) => {
    const ipFormat =
        /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(?:\/(?:3[0-2]|[1-2]?[0-9]))?$|^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.)?(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(?:\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){2}$/
    return v && !v.match(ipFormat) ? t('IP_FORMAT_REQUIRED') : true
}

const requiredRule: Validator = (v, t) => {
    return !!v || t('Required')
}

const numberRule: Validator = (v, t) => {
    return v && isNaN(Number(v)) ? t('NUMBER_REQUIRED') : true
}

const ipRule: Validator = (v, t) => {
    const ipFormat =
        /^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/
    return v && !v.match(ipFormat) ? t('IP_FORMAT_REQUIRED') : true
}

export {ipOrRangeRule, requiredRule, numberRule, ipRule}
