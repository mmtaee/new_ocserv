type Validator = (value: string, t: (key: string) => string) => true | string

const ipOrRangeRule: Validator = (v, t) => {
    const ipFormat =
        /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(?:\/(?:3[0-2]|[1-2]?[0-9]))?$|^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.)?(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(?:\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){2}$/
    return v && !v.match(ipFormat) ? t('IP_FORMAT_WITH_RANGE_REQUIRED') : true
}

const requiredRule: Validator = (v, t) => {
    return !!v || t('Required')
}

const numberRule: Validator = (v, t) => {
    return v && isNaN(Number(v)) ? t('NUMBER_REQUIRED') : true
}

const ipWithNetmaskRule: Validator = (v, t) => {
    const ipv4Segment = '(25[0-5]|2[0-4]\\d|1\\d{2}|[1-9]?\\d)'
    const ipv4 = `(${ipv4Segment}\\.){3}${ipv4Segment}`
    const pattern = new RegExp(`^${ipv4}/${ipv4}$`)

    return v && !pattern.test(v) ? t('IP_WITH_NETMASK_FORMAT_REQUIRED') : true
}

const ipRule: Validator = (v, t) => {
    const ipFormat =
        /^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/
    return v && !v.match(ipFormat) ? t('IP_FORMAT_REQUIRED') : true
}


const domainRule: Validator = (v, t) => {
    const domainRegex = /^(?!:\/\/)([a-zA-Z0-9-]{1,63}\.)+[a-zA-Z]{2,}$/;
    return v && !domainRegex.test(v)
        ? t('DOMAIN_FORMAT_REQUIRED')
        : true
}


export {ipOrRangeRule, requiredRule, numberRule, ipRule, ipWithNetmaskRule, domainRule}
