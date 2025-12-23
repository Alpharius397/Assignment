import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Link, useNavigate } from "react-router-dom";
import { LoginInput, type LoginInputType } from "@/zod/login";
import useLogin from "@/hooks/useLogin";
import { toast } from "sonner";
import { useLocation } from "react-router-dom";
import { useEffect } from "react";

export default function Login() {
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    const error = location.state?.error;
    if (error) toast.error(error);

    const success = location.state?.success;

    if (success) toast.success(success);
  }, [location.state]);

  const {
    control,
    handleSubmit,
    formState: { errors, isValid },
  } = useForm<LoginInputType>({
    resolver: zodResolver(LoginInput),
    mode: "onChange",
    defaultValues: {
      user_name: "",
      password: "",
    },
  });

  const onSuccess = () => {
    toast.success("Login Successful!");
    navigate("/");
  };

  const onError = (err: string) => {
    toast.error(`Login Failed! Error: ${err}`);
  };

  const onFailure = () => {
    toast.error("Something went wrong!");
  };

  const { mutate } = useLogin({
    onSuccess,
    onFailure,
    onError,
  });

  const submit = (data: LoginInputType) => {
    mutate(data);
  };

  return (
    <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm">
        <div className="flex flex-col gap-6">
          <Card>
            <CardHeader>
              <CardTitle>Login to your account</CardTitle>
              <CardDescription>
                Enter your username below to login to your account
              </CardDescription>
            </CardHeader>
            <CardContent>
              <form onSubmit={(event) => event.preventDefault()}>
                <FieldGroup>
                  <Controller
                    control={control}
                    name="user_name"
                    render={({ field: { onChange, value } }) => (
                      <Field>
                        <FieldLabel htmlFor="email">Username</FieldLabel>
                        <Input
                          id="user_name"
                          type="text"
                          placeholder="Username"
                          onChange={onChange}
                          value={value}
                          required
                        />
                        {errors.user_name && (
                          <p className="text-sm text-red-600 m-0 w-full text-left">
                            {errors.user_name.message}
                          </p>
                        )}
                      </Field>
                    )}
                  />

                  <Controller
                    control={control}
                    name="password"
                    render={({ field: { onChange, value } }) => (
                      <Field>
                        <FieldLabel htmlFor="password">Password</FieldLabel>
                        <Input
                          id="password"
                          type="password"
                          placeholder="Password"
                          required
                          onChange={onChange}
                          value={value}
                        />
                        {errors.password && (
                          <p className="text-sm text-red-600 m-0 w-full text-left">
                            {errors.password.message}
                          </p>
                        )}
                      </Field>
                    )}
                  />

                  <Field>
                    <Button
                      className={!isValid ? "" : "cursor-pointer"}
                      type="submit"
                      onClick={handleSubmit(submit)}
                      disabled={!isValid}
                    >
                      Login
                    </Button>
                    <FieldDescription className="text-center">
                      Don&apos;t have an account?{" "}
                      <Link to={"/register"}>Sign up</Link>
                    </FieldDescription>
                  </Field>
                </FieldGroup>
              </form>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
