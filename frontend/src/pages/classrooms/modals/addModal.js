import React from "react";
import { useForm } from "react-hook-form";
import { useAddClassroomMutation } from "../api/addClassroomMutation";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import { CloseSvg } from "../../../components/close-svg";

const AddClassroomModal = ({ isOpen, toggleModal }) => {
  const form = useForm();
  const {
    reset,
    register,
    handleSubmit,
    formState: { errors },
  } = form;

  const addClassroomMutation = useAddClassroomMutation();
  const onSubmit = (data) => {
    addClassroomMutation.mutateAsync(data);
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
            <h2>Введите данные</h2>
            <Button onClick={toggleModal} styleFeature="close">
              <CloseSvg />
            </Button>
          </div>
          <form
              onSubmit={handleSubmit(onSubmit)}
              className={styles["groups-form"]}
          >
            <label>Введите аудиторию</label>
            <input
                {...register("classroomId", { required: true })}
                placeholder="103(б)"
                className={styles["input"]}
            />
            {errors.classroomId && (
                <span className={styles["error"]}>Это поле обязательно</span>
            )}
            <label>Введите вместимость аудитории</label>
            <input
                type="number"
                {...register("capacity", { required: true, min: 1 })}
                placeholder="46"
                className={styles["input"]}
            />
            {errors.capacity && (
                <span className={styles["error"]}>
              {errors.capacity.type === "required"
                  ? "Это поле обязательно"
                  : "Значение должно быть больше 0"}
            </span>
            )}
            <Button type="submit">Отправить</Button>
          </form>
        </div>
      </div>
  );
};

export default AddClassroomModal;
