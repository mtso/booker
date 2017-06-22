const Browser = require('zombie')

Browser.localhost('booker.com', process.env.PORT || 3750)

describe('User signs in', function() {
  const browser = new Browser()

  before(function(done) {
    browser.visit('/', () => done())
  })

  describe('click to /signup', function() {
    before(function(done) {
      browser
        .click('a[href="/signup"]')
        .then(done)
    })

    it('should be successful', function() {
      browser.assert.success()
    })

    it('should go to /signup', function() {
      browser.assert.url({ pathname: '/signup' })
    })

    describe('enter credentials', function() {
      before(function(done) {
        browser
          .fill('username', 'testuser1')
          .fill('password', 'testpass1')
          .pressButton('Sign In', done)
      })

      it('should be successful', function() {
        browser.assert.success()
      })

      it('should redirect to root', function() {
        browser.assert.url({ pathname: '/' })
      })
    })
  })
})

describe('User signs in', function() {
  const browser = new Browser()

  before(function(done) {
    browser.visit('/', () => done())
  })

  describe('click to /signup', function() {
    before(function(done) {
      browser
        .click('a[href="/signup"]')
        .then(() => browser
          .fill('username', 'testuser1')
          .fill('password', 'testpass1')
          .pressButton('Sign In', done))
    })

    it('should be successful', function() {
      browser.assert.success()
    })

    it('should redirect to root', function() {
      browser.assert.url({ pathname: '/' })
    })

    describe('can log out', function() {
      before(function(done) {
        browser
          .pressButton('Log Out', done)
      })

      it('should be successful', function() {
        browser.assert.success()
      })

      it('should redirect to root', function() {
        browser.assert.url({ pathname: '/' })
      }) 
    })
  })
})
