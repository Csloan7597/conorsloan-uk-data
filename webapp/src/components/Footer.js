import React, { Component } from 'react';

const footerStyle = {
  position: 'absolute',
  bottom: 0,
  width: '100%',
  /* Set the fixed height of the footer here */
  height: '60px',
  backgroundColor: '#f5f5f5'
}

const floating = {
  marginTop: "15px",
  marginBottom: "15px",
  marginRight: "15px",
  marginLeft: "15px"
}

// Hard coded for now
export default () => {
  return (
    <div className="footer" style = {footerStyle}>
      <div className = "container">
        <div className="copywrite pull-left" style={floating}>Conor Sloan 2017  	&copy;</div>
        <div className="social-links pull-right" style={floating}>
          TODO ICONS
        </div>
      </div>
    </div>
  )
}
