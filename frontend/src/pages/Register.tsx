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
import { toast } from "sonner";
import { RegisterInput, type RegisterInputType } from "@/zod/register";
import useRegister from "@/hooks/useRegister";

export default function Register() {
  const navigate = useNavigate();

  const {
    control,
    handleSubmit,
    formState: { errors, isValid },
  } = useForm<RegisterInputType>({
    resolver: zodResolver(RegisterInput),
    mode: "onChange",
    defaultValues: {
      user_name: "",
      email: "",
      password: "",
      confirm_password: "",
      aadhar: "",
    },
  });

  const onSuccess = () => {
    toast.success("Register Successful!");
    navigate("/login");
  };

  const onError = (err: string) => {
    toast.error(`Register Failed! Error: ${err}`);
  };

  const onFailure = () => {
    toast.error("Something went wrong!");
  };

  const { mutate } = useRegister({
    onSuccess,
    onFailure,
    onError,
  });

  const submit = (data: RegisterInputType) => {
    mutate(data);
  };

  return (
    <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm">
        <div className="flex flex-col gap-6">
          <Card>
            <CardHeader>
              <CardTitle>Create an account</CardTitle>
              <CardDescription>
                Enter your information below to create your account{" "}
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
                        <FieldLabel htmlFor="user_name">Username</FieldLabel>
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
                    name="email"
                    render={({ field: { onChange, value } }) => (
                      <Field>
                        <FieldLabel htmlFor="email">Email Address</FieldLabel>
                        <Input
                          id="email"
                          type="email"
                          placeholder="Email"
                          required
                          onChange={onChange}
                          value={value}
                        />
                        {errors.email && (
                          <p className="text-sm text-red-600 m-0 w-full text-left">
                            {errors.email.message}
                          </p>
                        )}
                      </Field>
                    )}
                  />

                  <Controller
                    control={control}
                    name="aadhar"
                    render={({ field: { onChange, value } }) => (
                      <Field>
                        <FieldLabel htmlFor="aadhar">Aadhar ID</FieldLabel>
                        <Input
                          id="aadhar"
                          type="password"
                          placeholder="12-digit Aadhar ID"
                          required
                          onChange={onChange}
                          value={value}
                        />
                        {errors.aadhar && (
                          <p className="text-sm text-red-600 m-0 w-full text-left">
                            {errors.aadhar.message}
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

                  <Controller
                    control={control}
                    name="confirm_password"
                    render={({ field: { onChange, value } }) => (
                      <Field>
                        <FieldLabel htmlFor="confirm_password">
                          Confirm Password
                        </FieldLabel>
                        <Input
                          id="confirm_password"
                          type="password"
                          placeholder="Confirm Password"
                          required
                          onChange={onChange}
                          value={value}
                        />
                        <FieldDescription>
                          Ensure that password and confirm password matches
                        </FieldDescription>
                        {errors.confirm_password && (
                          <p className="text-sm text-red-600 m-0 w-full text-left">
                            {errors.confirm_password.message}
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
                      Register
                    </Button>
                    <FieldDescription className="text-center">
                      Already have an account?{" "}
                      <Link to={"/login"}>Sign now</Link>
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
