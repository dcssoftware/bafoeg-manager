export function convertStructTypes(input?: any): any {
    if (input === null || typeof input !== 'object') {
        return input; // no object, nothing to do
    }

    for (var key in input) {
        if (input.hasOwnProperty(key)) {
            var value = input[key];

            // Recognize Date by value format (2025-02-14T15:41:29.142776171Z)
            // if (typeof value === 'string' && /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{9}Z$/.test(value)) {
            if (typeof value === 'string' && /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}/.test(value)) {
                input[key] = new Date(value);
            } else if (typeof value === 'object' && value !== null) {
                // convert nested objects recursively
                input[key] = convertStructTypes(value);
            }
        }
    }
    return input;
}