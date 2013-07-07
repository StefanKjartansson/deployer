Backbone = require 'Backbone'
chai = require 'chai'
sinon = require 'sinon'
request = require "request"
projects = require '../application/projects.coffee'


expect = chai.expect
s = process.env.TEST_SERVER
attrs = null

describe 'Server', ->

  it 'should be able to contact the server', (done) ->
    request "http://#{s}/projects/", (err, res, body) ->
      expect(res.statusCode).to.equal(200)
      attrs = JSON.parse(body)
      done()


describe 'ProjectCollection', ->

  collection = undefined

  beforeEach ->
    collection = new projects.Collection
    collection.reset()

  it 'should be defined', ->
    expect(collection).to.not.equal(undefined)

  it 'should be able to list and filter', (done) ->

    Backbone.sync = (method, model) ->
      model.set(attrs)

    collection.fetch()

    expect(collection.where({name: "real_test"}).length).to.equal(1)

    done()
