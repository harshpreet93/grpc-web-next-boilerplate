import Link from 'next/link'
import Layout from '../components/Layout'
import { GreeterClient } from '../service_client_out/ServiceServiceClientPb';
import { HelloRequest } from '../service_client_out/service_pb';
import React, { Component } from 'react'

interface Props {}

interface State {
  message: String,
}

export default class IndexPage extends Component<Props, State> {
  async updateMessage() {
    console.log("starting to update message");
    var client = new GreeterClient("http://localhost:8080");
    var req = new HelloRequest();
    req.setName("harshpreet");
    var newMessage  = (await client.sayHello(req, {})).getMessage();
    this.setState( {message: newMessage}, () => {
      console.log("message is "+newMessage);
    });
  }

  constructor(props: Props) {
    super(props);

    this.state = {
      message: "blah"
    };
  }


  render() {
    return (
      <Layout title="Home | Next.js + TypeScript Example">
        <h1 onClick={() => this.updateMessage()}>{this.state.message} 👋 </h1>
        <p>
          <Link href="/about">
            <a>About</a>
          </Link>
        </p>
      </Layout>
    )
  }

}

