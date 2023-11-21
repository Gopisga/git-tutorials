import React from 'react'

const Footer = () => {
  const footerstyle ={
    position: "sticky",
    top: "100vh",
    width: "100%",
   border : "6px solid limegreen"
  }
  return (
    <footer className= "bg-dark text-light py-3"  style ={footerstyle} >
      <p className = "text-center">
      copyright&copy;MyTodoslist.com
      </p>
    </footer>
  )
  }

  export default Footer;
