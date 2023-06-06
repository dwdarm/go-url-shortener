

export const getLink = async (slug) => {
    const res = await fetch(`/api/links/${slug}`);
    if (!res.ok) {
        throw new Error();
    }

    const json = await res.json();

    return json;
}

export const createLink = async (data = {}) => {
    const res = await fetch('/api/links', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    });

    if (!res.ok) {
        throw Error();
    }

    const json = await res.json()

    return json;
}