function greeter(person) {
    return "Hello, " + person.firstName + " " + person.lastName;
}
var user = {
    firstName: "Hokuto",
    lastName: "Shinoda"
};
document.body.innerHTML = greeter(user);
