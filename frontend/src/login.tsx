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
  const { setPage, accessToken, setAccessToken, setUserID } = useApp()
  const { handleError } = useErrorHandler()
  const { handleNotif } = useNotifHandler()


  const [formData, setFormData] = useState<FormData>(emptyFormData)

  if (accessToken !== null) {
    handleError("Already logged in")
    setPage("userStats")
  }


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
        const userID = data.data._id
        setUserID(userID)
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

      <div className="flex flex-col items-start">
        <h1 className="font-bold text-8xl text-gray-50 mb-10">Login</h1>

        <h2 className="font-bold text-3xl text-gray-50">Username</h2>
        <input
          className="textInput"
          name="username"
          value={formData.username}
          onChange={handleChange}
          placeholder="Username"
        />

        <h2 className="font-bold text-3xl text-gray-50">Password</h2>
        <input
          className="textInput"
          type="password"
          name="password"
          value={formData.password}
          onChange={handleChange}
          placeholder="Password"
        />

        <button
          className="hover:bg-neutral-600 mt-3.5 text-2xl py-2 px-4 text-gray-50 border-neutral-400 bg-neutral-700 border-5"
          onClick={handleSubmit}
        >Submit</button>
      </div>

    </>
  );
}
export default Login;
