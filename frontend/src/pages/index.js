import React from "react"
import { useFirebase } from "gatsby-plugin-firebase"

import Layout from "../components/layout"
import SEO from "../components/seo"
import LoadingIndicator from "../components/loading"

const IndexPage = () => {
  const [message, setMessage] = React.useState(null)

  useFirebase(firebase => {
    firebase
      .firestore()
      .collection("babies")
      .doc("lukeandrebecca")
      .onSnapshot(
        snapshot => {
          setMessage(snapshot.data().status)
        },
        error => {
          setMessage("No Idea")
        }
      )
  }, [])

  return (
    <Layout>
      <SEO title="Home" />
      <div
        style={{
          width: "100%",
          height: "100",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <LoadingIndicator loading={!message} />
        <h1>{message ? message : ""}</h1>
      </div>
    </Layout>
  )
}

export default IndexPage
