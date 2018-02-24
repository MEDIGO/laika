import React from 'react'
import { string, func, node } from 'prop-types'

import Alert from './Alert'
import Button from './Button'

import './Form.css'

export default function Form({ submitText, onSubmit, errorText, children }) {
  const handleSubmit = e => {
    e.preventDefault()
    onSubmit(e)
  }

  return (
    <form onSubmit={handleSubmit}>
      {errorText ? <Alert type='danger'>{errorText}</Alert> : null}

      {children}
      <div className='lk-form__controls'>
        <Button type='primary' label={submitText} onClick={handleSubmit} />
      </div>
    </form>
  )
}

Form.propTypes = {
  submitText: string,
  onSubmit: func.isRequired,
  errorText: string,
  children: node
}

Form.defaultProps = {
  submitText: 'Submit',
  errorText: null,
  children: null
}
