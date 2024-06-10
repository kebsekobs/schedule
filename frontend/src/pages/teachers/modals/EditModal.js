import React from "react";
import { useForm } from "react-hook-form";
import { useEditTeacherMutation } from "../api/EditTeacherMutation";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import {CloseSvg} from "../../../components/close-svg";

const EditGroupModal = ({ original, isOpen, toggleModal, id }) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm({
    defaultValues: {
      name: original.name
    },
  });

  const editTeacherMutation = useEditTeacherMutation();
  const onSubmit = (data) => {
    data.id = id;
    editTeacherMutation.mutateAsync(data);
    toggleModal();
    reset();
  };

  if (!isOpen) {
    return null;
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
