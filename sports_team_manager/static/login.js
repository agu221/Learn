async function loginUser(emailOrUsername, password){
    try{
        const response = await fetch('/users/login',{
            method: 'POST',
            headers:{
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({emailOrUsername,password})
        }
        );

        if(!response.ok){
            throw new Error('Login failed');
        }

        const data = await response.json();
        localStorage.setItem('accessToken', data.accessToken);
        localStorage.setItem('refreshToken', data.refreshToken);
        window.location.href = "/user-dashboard";
    }
    catch(error){
        console.error('Error during login:', error.message);
        alert(`Login error: ${error.message}`);
    }
}