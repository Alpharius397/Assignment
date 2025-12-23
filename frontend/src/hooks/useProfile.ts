import Axios, { URL } from "@/axios";
import { ProfileData, type ProfileDataType } from "@/zod/profile";
import { useQuery } from "@tanstack/react-query";

async function getProfile() {
  const response = await Axios.get(URL.Profile);

  return await ProfileData.parseAsync(response.data);
}

export default function useProfile(
) {
  const { data, isError, isLoading, refetch, isFetching } =
    useQuery<ProfileDataType>({
      queryFn: getProfile,
      queryKey: ["get-profile"],
      retry: 2,
    });

  return { data, isError, isLoading, refetch, isFetching };
}
