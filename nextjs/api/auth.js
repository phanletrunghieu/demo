export function signup(user) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve(user)
        }, 2000);
    })
}