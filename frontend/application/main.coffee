_ = require 'underscore'
_.str = require 'underscore.string'
_.mixin(_.str.exports())
Backbone = require 'Backbone'

projects = require './projects.coffee'

conn = new WebSocket("ws://127.0.0.1:3999/ws")

conn.onopen = ->
  console.log "WebSocket connected"


class App extends Backbone.Router

  routes:
    "": "index"
    "project-detail/:id/": "projectView"

  index: ->
    collection = new projects.Collection
    view = new projects.CollectionView
      model: collection
    $("#page_body").html view.el

  projectView: (id) ->
    model = new projects.Model
      id: id
    view = new projects.DetailView
      model: model
    $(".detail-container").html view.el

$ ->
  console.log "Starting"

  conn.onclose = (evt) ->
    console.log "closed"
  conn.onmessage = (evt) ->
    console.log evt.data
  conn.onerror = (evt) ->
    console.log evt

  app = new App()
  Backbone.history.start()
