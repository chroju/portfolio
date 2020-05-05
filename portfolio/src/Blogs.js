import React from 'react';
import axios from 'axios';
import xml2js from 'xml2js';

const rssURL = "https://chroju.github.io/atom.xml"

class Blogs extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      entries: []
    }
  }
  componentDidMount() {
    axios.get(rssURL)
      .then((res) => {
        for (let i = 0; i < 3; i++) {
          const entries = this.state.entries;
          xml2js.parseString(res.data, (err, xml) => {
            console.log(xml)
            const newEntries = entries.concat({ url: xml.rss.channel[0].item[i].link[0], title: xml.rss.channel[0].item[i].title[0] });
            this.setState({ entries: newEntries });
          });
        }
      })
      .catch(console.error);
  }
  render() {
    var links = [];
    for (const i of this.state.entries) {
      links.push(<li><a href="{i.url}">{i.title}</a></li>);
    }
    return (links);
  }
}

export default Blogs;
