import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useTeacherByIdQuery } from "../api/getTeacherById";
import { useEditTeacherMutation } from "../api/EditTeacherMutation";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import {CloseSvg} from "../../../components/close-svg";

const EditGroupModal = ({ isOpen, toggleModal, id }) => {
  const teacherByIdQuery = useTeacherByIdQuery(id);
  const [isDataLoaded, setIsDataLoaded] = useState(false);
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm({
    defaultValues: {
      name: "",
      initials: "",
    },
  });

  const editTeacherMutation = useEditTeacherMutation();
  const onSubmit = (data) => {
    data.id = id;
    editTeacherMutation.mutateAsync(data);
    toggleModal();
    reset();
  };

  useEffect(() => {
    if (teacherByIdQuery.data) {
      reset({
        name: teacherByIdQuery.data.name,
        initials: teacherByIdQuery.data.initials,
      });
      setIsDataLoaded(true);
    }
  }, [teacherByIdQuery.data, reset]);

  if (!isOpen) {
    return null;
  }

  if (!isDataLoaded) {
    return <div>Loading...</div>;
  }

  return (
    <div className={styles["backdrop"]}>
      <div className={styles["modal"]}>
        <div className={styles["modal-header"]}>
          <h2>Измените данные</h2>
          <Button onClick={toggleModal} styleFeature="close">
            <CloseSvg />
          </Button>
        </div>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className={styles["groups-form"]}
        >
          <label>Измените ФИО</label>
          <input
            {...register("name", { required: true })}
            placeholder="Иванов И.И."
            className={styles["input"]}
          />
          {errors.name && (
            <span className={styles["error"]}>Это поле обязательно</span>
          )}
          <Button type="submit">Oтправить</Button>
        </form>
      </div>
    </div>
  );
};

export default EditGroupModal;
