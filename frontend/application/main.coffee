_ = require 'underscore'
_.str = require 'underscore.string'
_.mixin(_.str.exports())
Backbone = require 'Backbone'

projects = require './projects.coffee'


class App extends Backbone.Router

  routes:
    "": "index"
    "project-detail/:id/": "projectView"

  index: ->
    console.log "Index view"
    collection = new projects.Collection
    view = new projects.CollectionView
      model: collection
    $("#page_body").html view.el

  projectView: (id) ->
    console.log "Project view"


$ ->
  console.log "Starting"
  app = new App()
  Backbone.history.start()
