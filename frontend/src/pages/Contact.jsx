import { useState } from 'react';
import {useNavigate} from "react-router-dom"
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { Checkbox } from "@/components/ui/checkbox"
import { Button } from "@/components/ui/button"

export default function Contact() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('I would like to receive more information');
  const [agreement, setAgreement] = useState(false);
  const navigate=useNavigate()

  const handleSubmit = async () => {

    console.log(name, email, message, agreement)
    if (!name || !/^[a-zA-Z\s]+$/.test(name)) {
      toast.error('Please enter a valid name with only alphabets.');
      return;
    }

    if (!email || !/\S+@\S+\.\S+/.test(email)) {
      toast.error('Please enter a valid email.');
      return;
    }

    if (!message) {
      toast.error('Please enter a message.');
      return;
    }

    if (!agreement) {
      toast.error('Please agree to the Terms & Conditions.');
      return;
    }

    const res= await fetch('/api/v1/add-contact',{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, email, message }),
    })
    const data = await res.json();

    if (data.status=="fail") {
      toast.error(data.error);
    } else {
      toast.success('Submission successful!');
      setTimeout(()=>{
        navigate('/')
      },2000)
    }
  };

  return (
    <div key="1" className="mt-12 p-24 rounded-md space-y-8">
      <div className="space-y-2">
        <h2 className="text-3xl font-bold">Contact Us</h2>
        <p className="text-gray-500 dark:text-gray-400">
          Please fill the below form and we will get back to you as soon as possible.
        </p>
      </div>
      <div className="space-y-4">
        <div className="space-y-2">
          <Label className="text-gray-600 dark:text-gray-400 required" htmlFor="name">
            Name
          </Label>
          <Input
            className="border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800"
            id="name"
            placeholder="Enter your name"
            required
            onChange={e=>setName(e.target.value)}
          />
        </div>
        <div className="flex space-x-2 space-y-2 items-end">
          <div className="w-3/4 space-y-2">
            <Label className="text-gray-600 dark:text-gray-400 required" htmlFor="email">
              Email
            </Label>
            <Input
              className="border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800"
              id="email"
              placeholder="Enter your email"
              required
              onChange={e=>setEmail(e.target.value)}
            />
          </div>
        </div>
        <div className="space-y-2">
          <Label className="text-gray-600 dark:text-gray-400 required" htmlFor="message">
            Message
          </Label>
          <textarea
            className="border-2 border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 w-full p-2"
            defaultValue={message}
            id="message"
            placeholder="Type your message"
            required
            rows="4"
            onChange={e=>setMessage(e.target.value)}
          />
        </div>
        <div className="flex items-center space-x-2">
          <Checkbox className="text-gray-600 dark:text-gray-400" id="agreement" required onClick={()=>setAgreement(prev=>!prev)} />
          <Label className="text-sm font-normal text-gray-600 dark:text-gray-400" htmlFor="agreement">
            I agree to the&ensp;
            <button className="underline underline-offset-2 text-gray-600 dark:text-gray-400">
              Terms & Conditions
            </button>
          </Label>
        </div>
        <Button className="w-full bg-gray-300 hover:bg-gray-700 text-black hover:text-white" type="submit"
        onClick={handleSubmit}>
          Submit
        </Button>
        <ToastContainer />
      </div>
    </div>
  )
}

