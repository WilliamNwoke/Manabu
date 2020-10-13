const Parser = require('rss-parser')

export function handler(event, context, callback) {
  const FEEDS = [
    'https://codepen.io/posts/feed',
    'https://hnrss.org/frontpage',
    'https://github-trends.ryotarai/rss/github_trends_javascript_daily.rss'
  ]

  let parser = new Parser()
  const feedRequest = FEEDS.map(feed => {
    return parser.parseURL(feed)

  })

  Promise.all(feedRequest).then(response => {
    callback(null, {
      statusCode: 200,
      body: JSON.stringify(response)
    })
  })
}