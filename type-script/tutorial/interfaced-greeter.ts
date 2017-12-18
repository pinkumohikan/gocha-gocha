interface Person {
    firstName: string;
    lastName: string;
}

function greeter(person: Person) {
    return "Hello, " + person.firstName + " " + person.lastName;
}

let user = {
    firstName: "Hokuto",
    lastName: "Shinoda"
};

document.body.innerHTML = greeter(user);
