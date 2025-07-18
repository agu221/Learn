async function loginUser(email, password){
    try{
        const response = await fetch('/login',{
            method: 'POST',
            headers:{
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({email,password})
        }
        );

        if(!response.ok){
            throw new Error('Login failed');
        }

        const data = await response.json();
        // localStorage.setItem('accessToken', data.accessToken);
        // localStorage.setItem('refreshToken', data.refreshToken);
        console.log('Login sucessful');        
    }
    catch(error){
        console.error('Error during login:', error.message);
        alert(`Login error: ${error.message}`);
    }
}