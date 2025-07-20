async function registerUser(Username, FirstName, LastName, DOB, Email,PhoneNumber,Password, confirmPasword, Role, Gender){
    try{
        if (confirmPasword != Password){
            throw new Error('Passwords do not match!');
        }
        const response = await fetch('/users/register',{
            method: 'POST',
            headers:{
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({Username, FirstName, LastName, DOB, Email,PhoneNumber,Password, Role, Gender})
        }
        );

        if(!response.ok){
            throw new Error('Login failed');
        }

        const data = await response.json();
        window.location.href = "/user-dashboard";
    }
    catch(error){
        console.error('Error during registration:', error.message);
        alert(`Registeration error: ${error.message}`);
    }
}