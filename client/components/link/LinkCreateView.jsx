import { useState, useRef } from 'react';
import { useMutation } from '@tanstack/react-query';
import { createLink } from '../../api/links.api';

const BASE_URL = import.meta.env.VITE_BASE_URL || '';

function LinkCreateResult({ link }) {
    return (
        <div className="notification is-info is-light">
            <p className="title is-5">
                <a style={{ textDecoration: 'none' }} href={`${BASE_URL}/${link?.Slug}`} target="_blank">{`${BASE_URL}/${link?.Slug}`}</a>
            </p>
            <p className="subtitle is-6">{link.Href}</p>
            <p className="buttons">
                <a href={`${BASE_URL}/${link?.Slug}`} target="_blank" className="button">
                    <span className="icon is-small">
                        <i className="fas fa-share"></i>
                    </span>
                </a>
                <button className="button" onClick={() => {
                    navigator.clipboard.writeText(`${BASE_URL}/${link?.Slug}`);
                    alert('Link copied!');
                }}>
                    <span className="icon is-small">
                        <i className="fas fa-copy"></i>
                    </span>
                </button>
                <a href={link?.QrCode} download="qrcode.png" className="button">
                    <span className="icon is-small">
                        <i className="fas fa-qrcode"></i>
                    </span>
                </a>
            </p>
        </div>
    );
}

export default function LinkCreateView() {
    const [ link, setLink ] = useState(null);
    const inputRef = useRef(null);

    const linkMutation = useMutation({
        mutationFn: (params) => createLink(params),
        onSuccess: (result) => setLink(result?.data),
        onError: (err) => console.error(err),
    })

    const handleSubmit = async (event) => {
        event.preventDefault();

        if (!linkMutation.isLoading) {
            linkMutation.mutate({ href: inputRef.current.value });
        }
    }

    return (
        <>

        <form onSubmit={handleSubmit}>
            <div className="field has-addons has-addons-centered is-fullwidth mb-5">
                <div className="control is-expanded">
                    <input type="url" ref={inputRef} name="href" className="input" required placeholder="Paste your URL" />
                </div>
                <div className="control">
                    <button type="submit" className={`${linkMutation.isLoading && 'is-loading'} button is-link`}>Shorten</button>
                </div>
            </div>
        </form>

        { link && <LinkCreateResult link={link} /> }

        </>
    );
}