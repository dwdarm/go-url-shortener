import NavBar from "../components/commons/NavBar";
import LinkCreateView from "../components/link/LinkCreateView";

export default function Home() {
    return (
        <div>
            <NavBar/>
            <div className="section is-medium">
                <div className="container">
                    <h1 className="title is-2 is-size-3-mobile has-text-centered">Link Shortener</h1>
                    <h3 className="subtitle is-5 is-size-6-mobile has-text-centered">Convert your URLs to be shorter and easier to read</h3>
                    <div className="mb-6"></div>
                    <div style={{ maxWidth: 600, margin: '0 auto' }}>
                        <LinkCreateView/>
                    </div>
                </div>
            </div>
        </div>
    );
}