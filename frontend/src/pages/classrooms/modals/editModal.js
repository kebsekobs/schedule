import React from "react";
import { useForm } from "react-hook-form";
import { useEditClassroomMutation } from "../api/editClassroomMutation";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import {CloseSvg} from "../../../components/close-svg";
import {useClassroomsQuery} from "../api/getClassroomsQuery";

const EditGroupModal = ({ isOpen, toggleModal, original }) => {
  const data = useClassroomsQuery().data;
  const existingClassroomsIds = data.map(group => group.id) ?? []
console.log(existingClassroomsIds);
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm({
    defaultValues: {
      id: original.id,
      capacity: original.capacity
    },
  });

  const editClassroomMutation = useEditClassroomMutation();
  const onSubmit = (data) => {
    editClassroomMutation.mutateAsync(data);
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
          <label>Введите аудиторию</label>
          <input
              {...register("id", {
                required: "Это поле обязательно",
                validate: value => value=== original.id || !existingClassroomsIds.includes(value) || "Такой ID уже существует"
              })}
              placeholder="103(б)"
              className={styles["input"]}
          />
          {errors.id && (
              <span className={styles["error"]}>{errors.id.message}</span>)}
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

export default EditGroupModal;
