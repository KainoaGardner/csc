import API_URL from "./env.tsx"
import { useState } from "react"
import { useApp, useErrorHandler, useNotifHandler } from "./appContext/useApp.tsx"

type FormData = {
  username: string;
  email: string;
  password: string;
  passwordConfirm: string;
}

const emptyFormData = {
  username: "",
  email: "",
  password: "",
  passwordConfirm: "",
}

function Register() {
  const { setPage } = useApp()
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
    if (formData.username.length === 0 || formData.email.length === 0 || formData.password.length === 0 || formData.passwordConfirm.length === 0) {
      handleError("Cannot have empty inputs")
      setFormData(emptyFormData)
    } else if (formData.password !== formData.passwordConfirm) {
      handleError("Passwords do not match")
      setFormData(emptyFormData)
    } else {
      postUser()
    }
  }

  const postUser = async () => {
    const postData = {
      username: formData.username,
      email: formData.email,
      password: formData.password,
    }

    try {
      const response = await fetch(API_URL + "user", {
        method: "POST",
        headers: {
          "Content-Type": "application/json; charset=utf-8",
        },
        body: JSON.stringify(postData),
      });

      const data = await response.json();
      if (response.ok) {
        handleNotif("User Registered")
        setPage("login")
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
        <h1 className="font-bold text-8xl text-gray-50 mb-10">Register</h1>

        <h2 className="font-bold text-3xl text-gray-50">Username</h2>
        <input
          className="textInput"
          name="username"
          value={formData.username}
          onChange={handleChange}
          placeholder="Username"
        />

        <h2 className="font-bold text-3xl text-gray-50">Email</h2>
        <input
          className="textInput"
          name="email"
          value={formData.email}
          onChange={handleChange}
          placeholder="Email"
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

        <h2 className="font-bold text-3xl text-gray-50">Confirm Password</h2>
        <input
          className="textInput"
          type="password"
          name="passwordConfirm"
          value={formData.passwordConfirm}
          onChange={handleChange}
          placeholder="Confirm Password"
        />


        <button
          className="hover:bg-neutral-600 mt-3.5 text-2xl py-2 px-4 text-gray-50 border-neutral-400 bg-neutral-700 border-5"
          onClick={handleSubmit}
        >Submit</button>
      </div>

    </>
  );
}
export default Register;

