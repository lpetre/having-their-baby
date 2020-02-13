import React from "react"
import { useFirebase } from "gatsby-plugin-firebase"

import Layout from "../components/layout"
import SEO from "../components/seo"
import LoadingIndicator from "../components/loading"

const IndexPage = () => {
  const [message, setMessage] = React.useState(null)
  const [subdomain, setSubdomain] = React.useState(null)

  useFirebase(firebase => {
    if (!!subdomain) {
      firebase
        .firestore()
        .collection("babies")
        .doc(subdomain)
        .onSnapshot(
          snapshot => {
            setMessage(snapshot.data() ? snapshot.data().status : "Who?")
          },
          error => {
            setMessage("No Idea")
          }
        )
    }
  }, [subdomain])

  React.useEffect(() => {
    if (!subdomain && typeof window !== "undefined") {
      const hostname = window.location.hostname ?? "";
      if (hostname === "localhost" || hostname === "127.0.0.1") {
        return setSubdomain(process.env.GATSBY_HAVINGTHEIRBABY_SUBDOMAIN);
      }
      else {
        let parts = hostname.split(".")
        if (parts.length === 3) {
          setSubdomain(parts[0]);
        }
        else {
          setSubdomain("www")
        }
      }
    }
  }, [subdomain])


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
