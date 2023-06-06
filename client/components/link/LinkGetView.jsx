import { useParams, useNavigate  } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { getLink } from '../../api/links.api';

export default function LinkGetView() {
    const { slug } = useParams();
    const navigate = useNavigate();
    useQuery({
        queryKey: ['links', 'id'],
        queryFn: () => getLink(slug),
        retry: false,
        onSuccess: (result) => window.location.replace(result?.data?.Href),
        onError: () => navigate('/'),
    });

    return null;
}