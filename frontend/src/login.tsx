import API_URL from "./env.tsx"
import { useState } from "react"
import { useApp, useErrorHandler, useNotifHandler } from "./appContext/useApp.tsx"

type FormData = {
  username: string;
  password: string;
}

const emptyFormData = {
  username: "",
  password: "",
}


function Login() {
  const { setPage, setAccessToken } = useApp()
  const { handleError } = useErrorHandler()
  const { handleNotif } = useNotifHandler()


  const [formData, setFormData] = useState<FormData>(emptyFormData)

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }))
  }

  const handleSubmit = () => {
    if (formData.username.length === 0 || formData.password.length === 0) {
      handleError("Cannot have empty inputs")
      setFormData(emptyFormData)
    } else {
      postUserLogin()
    }
  }

  const postUserLogin = async () => {
    const postData = {
      username: formData.username,
      password: formData.password,
    }

    try {
      const response = await fetch(API_URL + "user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json; charset=utf-8",
        },
        body: JSON.stringify(postData),
      });

      const data = await response.json();
      if (response.ok) {
        handleNotif("Logged in")
        const accessToken = data.data.accessToken
        setAccessToken(accessToken)
        setPage("home")
      } else {
        handleError(data.error);
        setFormData(emptyFormData)
      }
    } catch (error) {
      console.log(error);
    }
  };


  return (
    <>
      <h1>Login</h1>
      <hr />

      <input
        name="username"
        value={formData.username}
        onChange={handleChange}
        placeholder="Username"
      />
      <input
        type="password"
        name="password"
        value={formData.password}
        onChange={handleChange}
        placeholder="Password"
      />

      <button
        onClick={handleSubmit}
      >Submit</button>

    </>
  );
}
export default Login;
