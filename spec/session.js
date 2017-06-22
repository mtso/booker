const Browser = require('zombie')

Browser.localhost('booker.com', process.env.PORT || 3750)

describe('user visits signin', function() {
  const browser = new Browser()
  
  before(function(done) {
    browser.visit('/', () => done())
  })

  describe('signin', function() {
    before(function(done) {
      browser
        .fill('username', 'testuser1')
        .fill('password', 'testpass1')
        .pressButton('Submit', () => done())
    })

    it('should be successful', function() {
      browser.assert.success()
    })
  })
})
