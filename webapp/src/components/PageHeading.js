import React from 'react';
import {PageHeader} from 'react-bootstrap';

const style = {
  verticalAlign: "top"
}

export default ({title, tagline}) => {
  return (
    <PageHeader style={style}>
      {title}<br/>
      <small>{tagline}</small>
    </PageHeader>
  )
}
