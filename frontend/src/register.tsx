import { useState } from "react"

type FormData = {
  username: string;
  email: string;
  password: string;
}

function Register({ handleError }: { handleError: (err: string) => void }) {
  const [formData, setFormData] = useState<FormData>({
    username: "",
    email: "",
    password: "",
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }))
  }

  const handleSubmit = () => {
    if (!checkValidSubmit(formData)) {
      handleError("Cannot submit empty input")
    } else {
      console.log(formData)
    }
  }


  return (
    <>
      <h1>Register</h1>
      <hr />

      <input
        name="username"
        value={formData.username}
        onChange={handleChange}
        placeholder="Username"
      />
      <input
        name="email"
        value={formData.email}
        onChange={handleChange}
        placeholder="Email"
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
export default Register;

function checkValidSubmit(formData: FormData): boolean {
  if (formData.username.length === 0 || formData.email.length === 0 || formData.password.length === 0) {
    return false
  }

  return true
}
