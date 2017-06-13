const expect = require('chai').expect
const request = require('superagent')
const hostname = process.env.HOSTNAME || 'http://localhost:3750'

describe('auth', function() {
  describe('POST signup', function() {
    it('validates application/json Content-Type header', function(done) {
      request
        .post(hostname + '/auth/signup')
        .send({
          username: 'test1',
          password: 'pass1',
        })
        .then(function(resp) {
          if (!resp.ok) {
            done(resp.err)
          }

          

          done()
        })
        .catch(done)
    })
  })
})
