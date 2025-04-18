function Authorization(): string {
    return "Bearer " + localStorage.getItem("token");
}

export {Authorization}