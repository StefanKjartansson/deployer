chai = require 'chai'
expect = chai.expect

sinon = require 'sinon'
request = require "request"


s = process.env.TEST_SERVER


describe 'Projects', ->

  it 'should be able to contact the server', (done) ->
    request "http://#{s}/projects/", (err, res, body) ->
      expect(res.statusCode).to.equal(200)
      done()


