const httpProxy = require('http-proxy')
const proxy = httpProxy.createProxyServer()
const API_URL = process.env.API_URL || 'http://localhost:8700/v1'

export default function(req, res, next) {
  proxy.web(req, res, {
    target: API_URL
  })
}