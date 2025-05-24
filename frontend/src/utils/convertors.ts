const bytesToGB = (bytes: number): string => {
    return (bytes / (1024 ** 3)).toFixed(6); // returns GB as a string with 6 decimal places
}


export {bytesToGB}
